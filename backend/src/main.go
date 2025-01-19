package main

import (
	"context"
	"log"
	"net/http"
	"os"

	handlers "github.com/lkhtk/workout-log/handlers"

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
	workoutsCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("workouts")
	usersCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	measurementsCollection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("measurements")
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
		authorized.GET("/workouts/:id", workoutsHandler.GetOneWorkout)
		authorized.POST("/workouts", workoutsHandler.NewWorkout)
		authorized.PUT("/workouts/:id", workoutsHandler.UpdateWorkout)
		authorized.DELETE("/workouts/:id", workoutsHandler.DeleteWorkout)

		authorized.GET("/measurements", measurementsHandler.GetAllMeasurements)
		authorized.GET("/measurement", measurementsHandler.GetLatestMeasurement)
		authorized.POST("/measurements", measurementsHandler.Create)
	}
	server.Use(cors.Default())
	server.Run("0.0.0.0:8000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
