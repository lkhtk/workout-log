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
var server *gin.Engine

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
	server = gin.Default()
	store := cookie.NewStore([]byte("4roiHJKLnmz,xJJ((2-____))"))

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
	router.GET("/auth/google/sigout", authHandler.DestroyUserSession)
	router.GET("/health", health)
	authorized := router.Group("/api")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.GET("/workouts", workoutsHandler.ListWorkouts)
		authorized.GET("/workouts/:id", workoutsHandler.GetOneWorkout)
		authorized.POST("/workouts", workoutsHandler.NewWorkout)
		authorized.PUT("/workouts/:id", workoutsHandler.UpdateWorkout)
		authorized.DELETE("/workouts/:id", workoutsHandler.DeleteWorkout)
	}
	server.Use(cors.Default())
	server.Run("0.0.0.0:8000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
