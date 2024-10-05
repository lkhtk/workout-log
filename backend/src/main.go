package main

import (
	"context"
	"log"
	"net/http"
	"os"

	handlers "github.com/lkhtk/workout-log/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client
var workoutsHandler *handlers.WorkoutsHandler

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DATSBASE")).Collection("workouts")
	workoutsHandler = handlers.NewWorkoutsHandler(ctx, collection)
}
func main() {
	router := gin.Default()
	router.GET("/workouts", workoutsHandler.ListWorkouts)
	router.GET("/workouts/:id", workoutsHandler.GetOneWorkout)
	router.POST("/workouts", workoutsHandler.NewWorkout)
	router.PUT("/workouts/:id", workoutsHandler.UpdateWorkout)
	router.DELETE("/workouts/:id", workoutsHandler.DeleteWorkout)
	router.Run("localhost:8000")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
