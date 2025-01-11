package handlers

import (
	"context"
	"errors"
	"log"
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

const maxLengthOfFileld = 150

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
	findOpt := options.Find()
	page, _ := strconv.Atoi(c.Query("page"))
	var perPage int64 = 3
	filter := getFilter(c)
	total, _ := handler.collection.CountDocuments(handler.ctx, filter)

	findOpt.SetSkip((int64(page) - 1) * perPage)
	findOpt.SetLimit(perPage)
	findOpt.SetSort(bson.D{{"publishedat", -1}})
	cur, err := handler.collection.Find(handler.ctx, filter, findOpt)
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

	c.JSON(http.StatusOK, gin.H{"data": workouts, "total": total, "page": page, "last_page": math.Ceil(float64(total/perPage) + 1)})
}

func getFilter(c *gin.Context) bson.M {
	filter := bson.M{}
	session := sessions.Default(c)
	email, exists := session.Get("email").(string)
	if !exists || email == "" {
		return filter
	}
	filter["user"] = email
	searchDate := c.Query("date")
	period := c.Query("period")
	location := time.Now().Location()

	if searchDate == "" {
		searchDate = time.Now().String()
	}
	if period == "" {
		period = "month"
	}
	if searchDate != "" && period != "" {
		parsedDate, err := time.Parse("2006-01-02", searchDate)
		if err == nil {
			switch period {
			case "day":
				startOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, location)
				endOfDay := startOfDay.AddDate(0, 0, 1)
				filter["publishedat"] = bson.M{
					"$gte": startOfDay,
					"$lt":  endOfDay,
				}
			case "month":
				startOfMonth := time.Date(parsedDate.Year(), parsedDate.Month(), 1, 0, 0, 0, 0, location)
				endOfMonth := startOfMonth.AddDate(0, 1, 0)
				filter["publishedat"] = bson.M{
					"$gte": startOfMonth,
					"$lt":  endOfMonth,
				}
			case "year":
				startOfYear := time.Date(parsedDate.Year(), time.January, 1, 0, 0, 0, 0, location)
				endOfYear := startOfYear.AddDate(1, 0, 0)
				filter["publishedat"] = bson.M{
					"$gte": startOfYear,
					"$lt":  endOfYear,
				}
			}
		}
	}

	return filter
}

func (handler *WorkoutsHandler) NewWorkout(c *gin.Context) {

	workout, err := validateUserData(c)
	if err != nil {
		log.Println(http.StatusBadRequest)
		return
	}
	removeEmptyRepsSets(&workout)
	workout.ID = primitive.NewObjectID()
	workout.PublishedAt = time.Now()
	_, err = handler.collection.InsertOne(handler.ctx, workout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new workout"})
		return
	}
	c.JSON(http.StatusCreated, workout)
}

func (handler *WorkoutsHandler) UpdateWorkout(c *gin.Context) {
	id := c.Param("id")
	workout, err := validateUserData(c)
	if err != nil {
		log.Println(http.StatusBadRequest)
		return
	}
	removeEmptyRepsSets(&workout)

	for _, ex := range workout.Workout.Exercises {
		if workout.Workout.SetsCount < len(ex.Sets) {
			workout.Workout.SetsCount = len(ex.Sets)
		}
	}

	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err = handler.collection.UpdateOne(handler.ctx, bson.M{"_id": objectId},
		bson.D{
			{"$set",
				bson.D{
					{"user", workout.User},
					{"muscle_group", workout.MuscleGroup},
					{"coach", workout.Coach},
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

func validateUserData(c *gin.Context) (models.Workout, error) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return workout, err
	}
	tooLargeMessage := "User data too large"
	for _, ex := range workout.Workout.Exercises {
		if len(ex.Name) > maxLengthOfFileld {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": tooLargeMessage})
			return workout, errors.New(tooLargeMessage)
		}
	}
	for _, cardio := range workout.Workout.Cardio {
		if len(cardio.Type) > maxLengthOfFileld {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": tooLargeMessage})
			return workout, errors.New(tooLargeMessage)
		}
	}

	if len(workout.MuscleGroup) > maxLengthOfFileld {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": tooLargeMessage})
		return workout, errors.New(tooLargeMessage)
	}

	session := sessions.Default(c)
	email := session.Get("email")
	userId := email.(string)
	userErr := "User undefined"

	if userId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": userErr})
		return workout, errors.New(userErr)
	}
	workout.User = userId
	return workout, nil
}

func removeEmptyRepsSets(workout *models.Workout) {
	filteredExercises := make([]struct {
		Name string `json:"name" bson:"name"`
		Sets []struct {
			Reps   *int     `json:"reps"`
			Weight *float32 `json:"weight"`
		} `json:"sets"`
	}, 0)

	for _, exercise := range workout.Workout.Exercises {
		filteredSets := make([]struct {
			Reps   *int     `json:"reps"`
			Weight *float32 `json:"weight"`
		}, 0)
		for _, set := range exercise.Sets {
			if set.Reps != nil && *set.Reps != 0 {
				filteredSets = append(filteredSets, set)
			}
		}

		if len(filteredSets) > 0 {
			exercise.Sets = filteredSets
			filteredExercises = append(filteredExercises, exercise)
		}
	}
	workout.Workout.Exercises = filteredExercises
}
