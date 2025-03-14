package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"

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
	authorized := router.Group("/api")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.GET("/workouts", workoutsHandler.ListWorkouts)
		authorized.GET("/workouts/average", workoutsHandler.AverageWeight)
		authorized.GET("/workouts/top5", workoutsHandler.Top5)
		authorized.GET("/workouts/:id", workoutsHandler.GetOneWorkout)
		authorized.POST("/workouts", workoutsHandler.NewWorkout)
		authorized.PUT("/workouts/:id", workoutsHandler.UpdateWorkout)
		authorized.DELETE("/workouts/:id", workoutsHandler.DeleteWorkout)

		authorized.GET("/measurements", measurementsHandler.GetAllMeasurements)
		authorized.GET("/measurement", measurementsHandler.GetLatestMeasurement)
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
