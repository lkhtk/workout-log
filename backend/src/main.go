package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	handlers "github.com/lkhtk/workout-log/handlers"
	"github.com/lkhtk/workout-log/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx                 context.Context
	err                 error
	client              *mongo.Client
	workoutsHandler     *handlers.MongoConnectionHandler
	measurementsHandler *handlers.MongoConnectionHandler
	authHandler         *handlers.AuthHandler
	server              *gin.Engine
)

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	database := client.Database(os.Getenv("MONGO_DATABASE"))

	workoutsCollection := database.Collection("workouts")
	usersCollection := database.Collection("users")
	measurementsCollection := database.Collection("measurements")

	workoutsHandler = handlers.NewWorkoutHandler(ctx, workoutsCollection)
	measurementsHandler = handlers.NewMeasurementsHandler(ctx, measurementsCollection)
	authHandler = handlers.NewAuthHandler(ctx, usersCollection)

}
func main() {
	server = gin.Default()
	c := client.Database(os.Getenv("MONGO_DATABASE")).Collection("sessions")

	store := mongodriver.NewStore(c, 3600, true, []byte(os.Getenv("SESSION_SECRET")))
	store.Options(sessions.Options{Path: "/", Domain: "localhost"})
	server.Use(sessions.Sessions("client_session", store))
	server.Use(ClearCorruptedSessionMiddleware())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	if os.Getenv("STAND") == "prod" {
		store.Options(sessions.Options{HttpOnly: true, SameSite: 4, Secure: true})
	} else {
		store.Options(sessions.Options{HttpOnly: false})
	}
	router := server.Group("/")
	router.POST("/auth/google/sigin", authHandler.GoogleAuthHandler)
	router.POST("/auth/google/sigout", authHandler.DestroyUserSession)
	router.GET("/health", health)
	router.GET("/sitemap.xml", sitemap)
	router.GET("/robots.txt", robots)
	router.GET("/.well-known/security.txt", security)
	authorized := router.Group("/api")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.GET("/workouts", workoutsHandler.ListWorkouts)
		authorized.GET("/workouts/average", workoutsHandler.AverageWeight)
		authorized.GET("/workouts/top5", workoutsHandler.Top5)
		authorized.POST("/workouts", workoutsHandler.NewWorkout)
		authorized.PUT("/workouts/:id", workoutsHandler.UpdateWorkout)
		authorized.DELETE("/workouts/:id", workoutsHandler.DeleteWorkout)

		authorized.GET("/measurements", measurementsHandler.GetAllMeasurements)
		authorized.GET("/measurements/last", measurementsHandler.GetLatestMeasurement)
		authorized.POST("/measurements", measurementsHandler.Create)

		authorized.GET("/user/export", exportAccountData)
		authorized.POST("/user/wipe", clearAccountData)
		authorized.DELETE("/user", deleteAccount)
	}
	server.Use(cors.Default())
	server.Run("0.0.0.0:8000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
func sitemap(c *gin.Context) {
	currentDate := time.Now().Format("2006-01-02")
	host := c.Request.Host
	sitemap := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://%s/about</loc>
        <lastmod>%s</lastmod>
    </url>
    <url>
        <loc>https://%s/features</loc>
        <lastmod>%s</lastmod>
    </url>
	<url>
        <loc>https://%s/faq</loc>
        <lastmod>%s</lastmod>
    </url>
    <url>
        <loc>https://%s/pricing</loc>
        <lastmod>%s</lastmod>
    </url>
</urlset>`, host, currentDate, host, currentDate, host, currentDate)
	c.Data(http.StatusOK, "application/xml", []byte(sitemap))
}

func security(c *gin.Context) {
	email := os.Getenv("SECURITY_EMAIL")
	expDate := time.Now().AddDate(100, 0, 0).Format("2006-01-02")
	host := c.Request.Host
	security := fmt.Sprintf(`Contact: mailto:%s
Expires: %sT00:00:00.000Z
Preferred-Languages: en, ru
Canonical: https://%s/.well-known/security.txt
`, email, expDate, host)
	c.Data(http.StatusOK, "text/plain", []byte(security))
}
func robots(c *gin.Context) {
	robots := `User-agent: *
Disallow: /
Allow: /features
Allow: /faq
Allow: /pricing
Allow: /about`
	c.Data(http.StatusOK, "text/plain", []byte(robots))
}

func exportAccountData(c *gin.Context) {
	zipBuffer := new(bytes.Buffer)
	zipWriter := utils.NewZipWriter(zipBuffer)
	if err := workoutsHandler.ExportWorkouts(c, zipWriter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process workouts"})
		return
	}
	if err := measurementsHandler.ExportMeasurements(c, zipWriter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process measurements"})
		return
	}
	if err := zipWriter.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize ZIP archive"})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=user_data.zip")
	c.Header("Content-Type", "application/zip")
	c.Data(http.StatusOK, "application/zip", zipBuffer.Bytes())
	c.JSON(http.StatusOK, "OK")
}

func clearAccountData(ctx *gin.Context) {
	if err := workoutsHandler.CleanupUserWorkouts(ctx); err != nil {
		log.Println("failed to fetch workouts: %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "workouts cleanup err"})
		return
	}
	if err := measurementsHandler.CleanupUserWorkouts(ctx); err != nil {
		log.Println("failed to fetch measurements: %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cleanup err"})
		return
	}
	ctx.JSON(http.StatusOK, nil)
	return
}

func deleteAccount(ctx *gin.Context) {
	if err := workoutsHandler.CleanupUserWorkouts(ctx); err != nil {
		log.Println("failed to fetch workouts: %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "workouts cleanup err"})
		return
	}
	if err := measurementsHandler.CleanupUserWorkouts(ctx); err != nil {
		log.Println("failed to fetch measurements: %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cleanup err"})
		return
	}
	if err := authHandler.DeleteCurrentUser(ctx); err != nil {
		log.Println("failed to fetch measurements: %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cleanup err"})
		return
	}
	ctx.JSON(http.StatusOK, nil)
	return
}

func ClearCorruptedSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered from session panic:", r)
				c.SetCookie("client_session", "", -1, "/", "", false, true)
			}
		}()
		_ = sessions.Default(c).Get("token")
		c.Next()
	}
}
