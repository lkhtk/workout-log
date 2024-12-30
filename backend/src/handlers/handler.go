package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutsHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewWorkoutsHandler(ctx context.Context, collection *mongo.Collection) *WorkoutsHandler {
	return &WorkoutsHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *WorkoutsHandler) ListWorkouts(c *gin.Context) {
	cur, err := handler.collection.Find(handler.ctx, getFilter(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)
	workouts := make([]models.Workout, 0)
	for cur.Next(handler.ctx) {
		var workout models.Workout
		cur.Decode(&workout)
		workouts = append(workouts, workout)
	}
	for i, wrk := range workouts {
		for _, ex := range wrk.Workout.Exercises {
			if workouts[i].Workout.SetsCount < len(ex.Sets) {
				workouts[i].Workout.SetsCount = len(ex.Sets)
			}

		}
	}

	c.JSON(http.StatusOK, workouts)
}

func getFilter(c *gin.Context) bson.M {
	session := sessions.Default(c)
	email := session.Get("email")
	userId := email.(string)
	if userId != "" {
		return bson.M{"user": userId}
	}
	return bson.M{}
}

func (handler *WorkoutsHandler) NewWorkout(c *gin.Context) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	email := session.Get("email")
	//todo search user id by email
	userId := email.(string)
	workout.ID = primitive.NewObjectID()
	workout.User = userId
	workout.PublishedAt = time.Now()
	_, err := handler.collection.InsertOne(handler.ctx, workout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new workout"})
		return
	}
	c.JSON(http.StatusCreated, workout)
}

func (handler *WorkoutsHandler) UpdateWorkout(c *gin.Context) {
	id := c.Param("id")
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, ex := range workout.Workout.Exercises {
		if workout.Workout.SetsCount < len(ex.Sets) {
			workout.Workout.SetsCount = len(ex.Sets)
		}
	}

	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := handler.collection.UpdateOne(handler.ctx, bson.M{"_id": objectId},
		bson.D{
			{"$set",
				bson.D{
					{"user", workout.User},
					{"muscle_group", workout.MuscleGroup},
					{"workout",
						bson.D{
							{"exercises", workout.Workout.Exercises},
							{"cardio", workout.Workout.Cardio},
						},
					},
				},
			},
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workout)
}
func (handler *WorkoutsHandler) GetOneWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	cur := handler.collection.FindOne(handler.ctx, bson.M{
		"_id": objectId,
	})
	var workout models.Workout
	err := cur.Decode(&workout)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, gin.H{"error": "no documents found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, ex := range workout.Workout.Exercises {
		if workout.Workout.SetsCount < len(ex.Sets) {
			workout.Workout.SetsCount = len(ex.Sets)
		}
	}

	c.JSON(http.StatusOK, workout)
}

func (handler *WorkoutsHandler) DeleteWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := handler.collection.DeleteOne(handler.ctx, bson.M{
		"_id": objectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workout has been deleted"})
}
