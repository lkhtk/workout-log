package handlers

import (
	"context"
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
	if page <= 0 {
		page = 1
	}
	filter := getFilter(c)
	total, _ := handler.collection.CountDocuments(handler.ctx, filter)

	findOpt.SetSkip((int64(page) - 1) * perPage)
	findOpt.SetLimit(perPage)
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

	c.JSON(http.StatusOK, gin.H{"data": workouts, "total": total, "page": page, "last_page": math.Ceil(float64(total / perPage))})
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
