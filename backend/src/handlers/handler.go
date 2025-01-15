package handlers

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxFieldLength = 150
const perPage int64 = 3

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

func getCurrentUser(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	user_id, ok := session.Get("user_id").(string)
	if !ok || user_id == "" {
		return "", errors.New("user is not authenticated")
	}
	return user_id, nil
}

func handleDBError(c *gin.Context, err error, notFoundMessage string) {
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (handler *WorkoutsHandler) ListWorkouts(c *gin.Context) {
	userID, err := getCurrentUser(c)
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

func getFilter(c *gin.Context, userId string) bson.M {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return bson.M{}
	}
	filter := bson.M{
		"user_id": objectId,
	}
	period := c.Query("period")
	currentDate := time.Now().Truncate(24 * time.Hour)
	if period == "" {
		period = "all"
	}

	switch period {
	case "day":
		end := currentDate.AddDate(0, 0, 1)
		filter["publishedat"] = bson.M{
			"$gte": currentDate,
			"$lt":  end,
		}
	case "month":
		start := time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 1, 0)
		filter["publishedat"] = bson.M{
			"$gte": start,
			"$lt":  end,
		}
	case "year":
		start := time.Date(currentDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(1, 0, 0)
		filter["publishedat"] = bson.M{
			"$gte": start,
			"$lt":  end,
		}
	case "all":
	}
	return filter
}

func validateAndPrepareWorkout(c *gin.Context) (models.Workout, error) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		return workout, errors.New("invalid workout data")
	}
	workout.ID = primitive.NewObjectID()
	workout.PublishedAt = time.Now()
	userID, err := getCurrentUser(c)
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

func (handler *WorkoutsHandler) NewWorkout(c *gin.Context) {
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

func (handler *WorkoutsHandler) DeleteWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	userID, err := getCurrentUser(c)
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

func (handler *WorkoutsHandler) GetOneWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	userID, err := getCurrentUser(c)
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

func (handler *WorkoutsHandler) UpdateWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	userID, err := getCurrentUser(c)
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
