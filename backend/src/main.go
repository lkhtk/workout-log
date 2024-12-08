package main

import (
	"context"
	"log"
	"net/http"
	"os"

	handlers "github.com/lkhtk/workout-log/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var workoutsHandler *handlers.WorkoutsHandler
var authHandler *handlers.AuthHandler

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("workouts")
	workoutsHandler = handlers.NewWorkoutsHandler(ctx, collection)
	authHandler = handlers.NewAuthHandler(ctx, collection)

}
func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		router.GET("/workouts", workoutsHandler.ListWorkouts)
		router.GET("/workouts/:id", workoutsHandler.GetOneWorkout)
		router.POST("/workouts", workoutsHandler.NewWorkout)
		router.PUT("/workouts/:id", workoutsHandler.UpdateWorkout)
		router.DELETE("/workouts/:id", workoutsHandler.DeleteWorkout)
	}

	router.POST("/auth/google/sigin", authHandler.GoogleAuthHandler)
	router.GET("/auth/google/sigout", authHandler.DestroyUserSession)
	router.GET("/health", health)
	router.Use(cors.Default())
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*://localhost:*/*"}
	// config.AllowHeaders = []string{"Access-Control-Allow-Headers", "X-Requested-With,content-type"}
	// config.AllowHeaders = []string{"Origin"}
	// config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}

	// config.AllowCredentials = true
	// router.Use(cors.New(config))
	router.Run("0.0.0.0:8000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
