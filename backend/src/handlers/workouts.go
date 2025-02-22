package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"github.com/lkhtk/workout-log/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxFieldLength = 150
const minFieldLength = 3
const pageSize int64 = 4
const maxSetsCount = 10
const maxExsCount = 10

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
	_, userID, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	total, err := handler.collection.CountDocuments(handler.ctx, userID)
	if err != nil {
		handleDBError(c, err, "error counting documents")
		return
	}
	size, _ := strconv.ParseInt(c.Query("size"), 10, 64)
	if size == 0 {
		size = pageSize
	}

	findOptions := options.Find().
		SetSkip((int64(page) - 1) * size).
		SetLimit(size).
		SetSort(bson.D{{"publishedat", -1}})

	cur, err := handler.collection.Find(handler.ctx, userID, findOptions)
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
		"last_page": math.Ceil(float64(total) / float64(pageSize)),
	})
}

func validateAndPrepareWorkout(c *gin.Context) (models.Workout, error) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		return workout, errors.New("invalid workout data")
	}
	workout.ID = primitive.NewObjectID()
	workout.PublishedAt = time.Now()
	userIDPrimitive, _, err := getCurrentUser(c)
	if err != nil {
		return workout, err
	}
	workout.UserID = userIDPrimitive
	if len(workout.MuscleGroup) > maxFieldLength || len(workout.MuscleGroup) < minFieldLength {
		return workout, errors.New("Invalid MuscleGroup:" + workout.MuscleGroup)
	}
	if len(workout.Workout.Exercises) > maxExsCount || len(workout.Workout.Cardio) > maxExsCount {
		return workout, errors.New("too much Exercises")
	}
	for _, ex := range workout.Workout.Exercises {
		if len(ex.Name) > maxFieldLength || len(ex.Name) < minFieldLength {
			return workout, errors.New("Invalid exercise name:" + ex.Name)
		}
		if len(ex.Sets) > maxSetsCount {
			return workout, errors.New("too much sets:" + ex.Name)
		}
		for _, set := range ex.Sets {
			if set.Reps == nil || *set.Reps <= 0 {
				return workout, errors.New("Invalid workout data:" + ex.Name)
			}
		}
	}
	for _, cardio := range workout.Workout.Cardio {
		if len(cardio.Type) > maxFieldLength || len(cardio.Type) < minFieldLength {
			return workout, errors.New("Invalid cardio name:" + cardio.Type)
		}
		if cardio.Time <= 0 {
			return workout, errors.New("Cardio time should not be zero:" + cardio.Type)
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
	_, userID, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	result := handler.collection.FindOne(handler.ctx, bson.M{"_id": objectId, "user_id": userID})
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
	_, userID, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var workout models.Workout
	err = handler.collection.FindOne(handler.ctx, bson.M{"_id": objectId, "user_id": userID}).Decode(&workout)
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
	_, userID, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var oldWorkout models.Workout
	err = handler.collection.FindOne(handler.ctx, bson.M{"_id": objectId, "user_id": userID}).Decode(&oldWorkout)
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

func (handler *MongoConnectionHandler) ExportWorkouts(ctx *gin.Context, zipWriter *utils.ZipWriter) error {
	_, userID, err := getCurrentUser(ctx)
	if err != nil {
		return err
	}
	cursor, err := handler.collection.Find(ctx, userID)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return err
	}

	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return zipWriter.AddFile("workouts.json", data)
}

func (handler *MongoConnectionHandler) CleanupUserWorkouts(c *gin.Context) error {
	_, userID, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	_, err = handler.collection.DeleteMany(c, userID)
	return err
}
