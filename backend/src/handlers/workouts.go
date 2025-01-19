package handlers

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxFieldLength = 150
const perPage int64 = 3

type MongoConnectionHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewWorkoutHandler(ctx context.Context, collection *mongo.Collection) *MongoConnectionHandler {
	return &MongoConnectionHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *MongoConnectionHandler) ListWorkouts(c *gin.Context) {
	userID, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	filter := getFilter(c, userID)
	total, err := handler.collection.CountDocuments(handler.ctx, filter)
	if err != nil {
		handleDBError(c, err, "error counting documents")
		return
	}

	findOptions := options.Find().
		SetSkip((int64(page) - 1) * perPage).
		SetLimit(perPage).
		SetSort(bson.D{{"publishedat", -1}})

	cur, err := handler.collection.Find(handler.ctx, filter, findOptions)
	if err != nil {
		handleDBError(c, err, "error fetching workouts")
		return
	}
	defer cur.Close(handler.ctx)

	var workouts []models.Workout
	if err := cur.All(handler.ctx, &workouts); err != nil {
		handleDBError(c, err, "error decoding workouts")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      workouts,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total) / float64(perPage)),
	})
}

func validateAndPrepareWorkout(c *gin.Context) (models.Workout, error) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		return workout, errors.New("invalid workout data")
	}
	workout.ID = primitive.NewObjectID()
	workout.PublishedAt = time.Now()
	userID, err := GetCurrentUser(c)
	if err != nil {
		return workout, err
	}
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return workout, err
	}
	workout.UserID = uid
	for _, ex := range workout.Workout.Exercises {
		if len(ex.Name) > maxFieldLength {
			return workout, errors.New("exercise name too long")
		}
		for _, set := range ex.Sets {
			if set.Reps != nil && *set.Reps == 0 {
				set.Reps = nil
			}
		}
	}
	return workout, nil
}

func (handler *MongoConnectionHandler) NewWorkout(c *gin.Context) {
	workout, err := validateAndPrepareWorkout(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = handler.collection.InsertOne(handler.ctx, workout)
	if err != nil {
		handleDBError(c, err, "error inserting workout")
		return
	}
	c.JSON(http.StatusCreated, workout)
}

func (handler *MongoConnectionHandler) DeleteWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userObject, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	result := handler.collection.FindOne(handler.ctx, bson.M{"_id": objectId, "user_id": userObject})
	if err := result.Err(); err != nil {
		handleDBError(c, err, "workout not found")
		return
	}

	_, err = handler.collection.DeleteOne(handler.ctx, bson.M{"_id": objectId})
	if err != nil {
		handleDBError(c, err, "error deleting workout")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout deleted successfully"})
}

func (handler *MongoConnectionHandler) GetOneWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userObject, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var workout models.Workout
	err = handler.collection.FindOne(handler.ctx, bson.M{"_id": objectId, "user_id": userObject}).Decode(&workout)
	if err != nil {
		handleDBError(c, err, "workout not found")
		return
	}
	c.JSON(http.StatusOK, workout)
}

func (handler *MongoConnectionHandler) UpdateWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userObject, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var oldWorkout models.Workout
	err = handler.collection.FindOne(handler.ctx, bson.M{"_id": objectId, "user_id": userObject}).Decode(&oldWorkout)
	if err != nil {
		handleDBError(c, err, "workout not found")
		return
	}
	newWorkout, err := validateAndPrepareWorkout(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{
		"$set": bson.M{
			"user":         newWorkout.UserID,
			"muscle_group": newWorkout.MuscleGroup,
			"coach":        newWorkout.Coach,
			"workout": bson.M{
				"exercises": newWorkout.Workout.Exercises,
				"cardio":    newWorkout.Workout.Cardio,
			},
		},
	}
	_, err = handler.collection.UpdateOne(handler.ctx, bson.M{"_id": objectId}, update)
	if err != nil {
		handleDBError(c, err, "error updating workout")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout updated successfully", "workout": newWorkout})
}
