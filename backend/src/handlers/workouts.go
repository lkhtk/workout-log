package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
const minIntValue = 0
const minFloatValue = 0.1

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
	_, userIdAsFilter, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil || page <= 0 {
		page = 1
	}

	size, err := strconv.ParseInt(c.Query("size"), 10, 64)
	if err != nil || size <= 0 {
		size = pageSize
	}

	var total int64

	total, err = handler.collection.CountDocuments(handler.ctx, userIdAsFilter)
	if err != nil {
		handleDBError(c, err, "error counting documents")
		return
	}

	findOptions := options.Find().
		SetSkip((page - 1) * size).
		SetLimit(size).
		SetSort(bson.D{{Key: "publishedat", Value: -1}})

	cur, err := handler.collection.Find(handler.ctx, userIdAsFilter, findOptions)
	if err != nil {
		handleDBError(c, err, "error fetching workouts")
		return
	}
	defer cur.Close(handler.ctx)

	workouts := make([]models.Workout, 0, size)
	if err := cur.All(handler.ctx, &workouts); err != nil {
		handleDBError(c, err, "error decoding workouts")
		return
	}

	lastPage := int64(1)
	if total > 0 {
		lastPage = (total + size - 1) / size
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      workouts,
		"total":     total,
		"page":      page,
		"last_page": lastPage,
	})
}

func validateAndPrepareWorkout(c *gin.Context) (*models.Workout, error) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		return nil, errors.New("invalid workout data")
	}
	workout.ID = primitive.NewObjectID()
	workout.PublishedAt = time.Now()
	userIDPrimitive, _, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	workout.UserID = userIDPrimitive
	if err := validateWorkoutFields(&workout); err != nil {
		return nil, err
	}
	return &workout, nil
}

func validateWorkoutFields(workout *models.Workout) error {
	var errs []error
	if len(workout.MuscleGroup) > maxFieldLength || len(workout.MuscleGroup) < minFieldLength {
		errs = append(errs, errors.New("Workout: invalid MuscleGroup"))
	}

	if len(workout.Workout.Exercises) <= minIntValue && len(workout.Workout.Cardio) <= minIntValue {
		errs = append(errs, errors.New("Workout: exercises or cardio required"))
	}

	if len(workout.Workout.Exercises) > maxExsCount || len(workout.Workout.Cardio) > maxExsCount {
		errs = append(errs, errors.New("Workout: too many exercises or cardio workouts"))
	}

	if err := validateReview(workout.Review); err != nil {
		errs = append(errs, err)
	}

	for _, ex := range workout.Workout.Exercises {
		if err := validateExercise(ex); err != nil {
			errs = append(errs, err)
		}
	}
	for _, cardio := range workout.Workout.Cardio {
		if err := validateCardio(cardio); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func validateReview(rev models.Review) error {
	const maxDurationMin = 480
	const maxSleepHrs = 24
	const maxIntensity = 10
	var errs []string
	if (models.Review{}) == rev {
		return errors.New("Review is required")
	}
	if rev.PerceivedIntensity == nil {
		errs = append(errs, "Review: invalid is required")
	} else if *rev.PerceivedIntensity < minIntValue || *rev.PerceivedIntensity > maxIntensity {
		errs = append(errs, fmt.Sprintf("Review: invalid intensity: %d", *rev.PerceivedIntensity))
	}
	if rev.Duration == nil {
		errs = append(errs, "Review: duration is required")
	} else if *rev.Duration < minIntValue || *rev.Duration > maxDurationMin {
		errs = append(errs, fmt.Sprintf("Review: invalid duration: %d", *rev.Duration))
	}
	if rev.HrsSlept == nil {
		errs = append(errs, "Review: hrs is required")
	} else if *rev.HrsSlept > maxSleepHrs || *rev.HrsSlept < minIntValue {
		errs = append(errs, "Review: invalid sleep hours")
	}
	if len(errs) > 0 {
		return fmt.Errorf("Review validation errors: %s", strings.Join(errs, "; "))
	}
	return nil
}

func validateExercise(ex models.Exercise) error {
	var errs []string
	minSet := 1
	if len(ex.Name) > maxFieldLength || len(ex.Name) < minFieldLength {
		errs = append(errs, fmt.Sprintf("invalid exercise name: %q", ex.Name))
	}
	if len(ex.Sets) > maxSetsCount || len(ex.Sets) < minSet {
		errs = append(errs, fmt.Sprintf("invalid sets in exercise: %q", ex.Name))
	}
	for _, set := range ex.Sets {
		if set.Reps == nil || *set.Reps <= minIntValue {
			errs = append(errs, fmt.Sprintf("invalid reps in exercise: %q", ex.Name))
		}
		if set.Weight == nil || *set.Weight < minIntValue {
			errs = append(errs, fmt.Sprintf("invalid weight in exercise: %q", ex.Name))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("Exercise validation errors: %s", strings.Join(errs, "; "))
	}
	return nil
}

func validateCardio(cardio models.Cardio) error {
	var errs []string
	const (
		minHeart     = 60
		maxHeartRate = 220
		maxSpeed     = 200.0
	)
	if len(cardio.Type) < minFieldLength || len(cardio.Type) > maxFieldLength {
		errs = append(errs, fmt.Sprintf("invalid cardio type length: %q", cardio.Type))
	}
	if cardio.Time < minIntValue {
		errs = append(errs, fmt.Sprintf("cardio time must be >= %d", minIntValue))
	}
	if cardio.HeartRate < minHeart {
		errs = append(errs, fmt.Sprintf("heart rate too low: %d", cardio.HeartRate))
	} else if cardio.HeartRate > maxHeartRate {
		errs = append(errs, fmt.Sprintf("heart rate too high: %d", cardio.HeartRate))
	}
	if cardio.Distance < minFloatValue {
		errs = append(errs, fmt.Sprintf("distance too low: %.2f", cardio.Distance))
	}
	if cardio.Speed < minFloatValue {
		errs = append(errs, fmt.Sprintf("speed too low: %.2f", cardio.Speed))
	} else if cardio.Speed >= maxSpeed {
		errs = append(errs, fmt.Sprintf("speed too high: %.2f", cardio.Speed))
	}
	if cardio.Calories < 0 {
		errs = append(errs, fmt.Sprintf("calories cannot be negative: %d", cardio.Calories))
	}
	if len(errs) > 0 {
		return fmt.Errorf("cardio validation errors: %s", strings.Join(errs, "; "))
	}
	return nil
}

func (handler *MongoConnectionHandler) NewWorkout(c *gin.Context) {
	workout, err := validateAndPrepareWorkout(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := handler.collection.InsertOne(handler.ctx, workout); err != nil {
		handleDBError(c, err, "error inserting workout")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":          workout.ID.Hex(),
		"publishedAt": workout.PublishedAt,
	})
}

func (handler *MongoConnectionHandler) DeleteWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	userID, _, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	result := handler.collection.FindOneAndDelete(handler.ctx, bson.M{"_id": objectId, "user_id": userID})
	if err := result.Err(); err != nil {
		handleDBError(c, err, "workout not found or already deleted")
		return
	}

	c.JSON(http.StatusOK, gin.H{"acknowledged": true})
}

func (handler *MongoConnectionHandler) UpdateWorkout(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID format"})
		return
	}

	userID, _, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newWorkout, err := validateAndPrepareWorkout(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"muscle_group":      newWorkout.MuscleGroup,
			"coach":             newWorkout.Coach,
			"review":            newWorkout.Review,
			"workout.exercises": newWorkout.Workout.Exercises,
			"workout.cardio":    newWorkout.Workout.Cardio,
		},
	}

	result := handler.collection.FindOneAndUpdate(
		handler.ctx,
		bson.M{"_id": objectId, "user_id": userID},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if err := result.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout not found or update failed"})
		return
	}

	var updatedWorkout models.Workout
	if err := result.Decode(&updatedWorkout); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding updated workout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Workout updated successfully",
		"workout":   updatedWorkout,
		"updatedAt": time.Now(),
	})
}

func (handler *MongoConnectionHandler) ExportWorkouts(ctx *gin.Context, zipWriter *utils.ZipWriter) error {
	_, userID, err := getCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return err
	}
	opts := options.Find().SetProjection(bson.D{{"user_id", 0}, {"_id", 0}, {"user", 0}, {"sets_count", 0}})
	cursor, err := handler.collection.Find(ctx, userID, opts)
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return err
	}
	_, err = handler.collection.DeleteMany(c, userID)
	return err
}
func (handler *MongoConnectionHandler) Stats(c *gin.Context) {
	_, userIdAsFilter, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	period := c.DefaultQuery("size", "all") // all | year | month | week

	var matchFilter bson.D
	switch period {
	case "month-12":
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", time.Now().AddDate(-1, 0, 0)}}}}
	case "month-6":
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", time.Now().AddDate(0, -6, 0)}}}}
	case "month-3":
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", time.Now().AddDate(0, -3, 0)}}}}
	case "month-1":
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", time.Now().AddDate(0, -1, 0)}}}}
	case "week":
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", time.Now().AddDate(0, 0, -7)}}}}
	default:
		matchFilter = bson.D{}
	}

	pipeline := bson.A{}
	if len(matchFilter) > 0 {
		pipeline = append(pipeline, bson.D{{"$match", matchFilter}})
	}

	pipeline = append(pipeline,
		bson.D{{"$match", userIdAsFilter}},
		bson.D{{"$sort", bson.D{{"publishedat", -1}}}},
		bson.D{{"$project", bson.D{
			{"muscle_group", 1},
			{"publishedat", 1},
			{"coach", 1},
			{"review", 1},
			{"exercises", bson.D{
				{"$map", bson.D{
					{"input", "$workout.exercises"},
					{"as", "exercise"},
					{"in", bson.D{
						{"name", "$$exercise.name"},
						{"average_weight_per_exercise", bson.D{
							{"$cond", bson.D{
								{"if", bson.D{
									{"$gt", bson.A{
										bson.D{{"$sum", bson.D{
											{"$map", bson.D{
												{"input", "$$exercise.sets"},
												{"as", "set"},
												{"in", "$$set.reps"},
											}},
										}}},
										0,
									}},
								}},
								{"then", bson.D{
									{"$divide", bson.A{
										bson.D{{"$sum", bson.D{
											{"$map", bson.D{
												{"input", "$$exercise.sets"},
												{"as", "set"},
												{"in", bson.D{
													{"$multiply", bson.A{"$$set.weight", "$$set.reps"}},
												}},
											}},
										}}},
										bson.D{{"$sum", bson.D{
											{"$map", bson.D{
												{"input", "$$exercise.sets"},
												{"as", "set"},
												{"in", "$$set.reps"},
											}},
										}}},
									}},
								}},
								{"else", 0},
							}},
						}},
					}},
				}},
			}},
		}}})

	opts := options.Aggregate().SetBatchSize(500)

	cur, err := handler.collection.Aggregate(handler.ctx, pipeline, opts)
	if err != nil {
		handleDBError(c, err, "error fetching average weights")
		return
	}
	defer cur.Close(handler.ctx)

	var workouts []models.Stats
	if err := cur.All(handler.ctx, &workouts); err != nil {
		handleDBError(c, err, "error decoding workouts")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": workouts})
}

func (handler *MongoConnectionHandler) Top5(c *gin.Context) {
	const limitOfTopEx int64 = 5
	_, userID, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	period := c.DefaultQuery("size", "all") // 'all', 'month', 'week'

	var matchFilter bson.D
	switch period {
	case "month":
		oneMonthAgo := time.Now().AddDate(0, -1, 0)
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", oneMonthAgo}}}}
	case "week":
		oneWeekAgo := time.Now().AddDate(0, 0, -7)
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", oneWeekAgo}}}}
	default:
		matchFilter = bson.D{}
	}
	pipeline := bson.A{}
	if len(matchFilter) > 0 {
		pipeline = append(pipeline, bson.D{{"$match", matchFilter}})
	}

	pipeline = append(pipeline,
		bson.D{{"$match", userID}},
		bson.D{{"$sort", bson.D{{"publishedat", -1}}}},
		bson.D{{"$unwind", "$workout.exercises"}},
		bson.D{{"$group", bson.D{
			{"_id", "$workout.exercises.name"},
			{"count", bson.D{{"$sum", 1}}},
		}}},
		bson.D{{"$sort", bson.D{{"count", -1}}}},
		bson.D{{"$limit", limitOfTopEx}},
	)

	cur, err := handler.collection.Aggregate(c, pipeline)
	if err != nil {
		handleDBError(c, err, "failed to execute aggregation")
		return
	}
	defer cur.Close(handler.ctx)

	var workouts []models.Top
	if err := cur.All(handler.ctx, &workouts); err != nil {
		handleDBError(c, err, "error decoding workouts")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": workouts})
}
