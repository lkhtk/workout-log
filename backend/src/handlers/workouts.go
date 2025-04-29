package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
		errs = append(errs, fmt.Errorf("invalid MuscleGroup: %q", workout.MuscleGroup))
	}

	if len(workout.Workout.Exercises) > maxExsCount || len(workout.Workout.Cardio) > maxExsCount {
		errs = append(errs, errors.New("too many exercises or cardio workouts"))
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

func validateExercise(ex models.Exercise) error {
	if len(ex.Name) > maxFieldLength || len(ex.Name) < minFieldLength {
		return fmt.Errorf("invalid exercise name: %q", ex.Name)
	}
	if len(ex.Sets) > maxSetsCount {
		return fmt.Errorf("too many sets in exercise: %q", ex.Name)
	}
	for _, set := range ex.Sets {
		if set.Reps == nil || *set.Reps <= 0 {
			return fmt.Errorf("invalid reps in exercise: %q", ex.Name)
		}
	}
	return nil
}

func validateCardio(cardio models.Cardio) error {
	if len(cardio.Type) > maxFieldLength || len(cardio.Type) < minFieldLength {
		return fmt.Errorf("invalid cardio name: %q", cardio.Type)
	}
	if cardio.Time <= 0 {
		return fmt.Errorf("cardio time should be greater than zero: %q", cardio.Type)
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
func (handler *MongoConnectionHandler) AverageWeight(c *gin.Context) {
	_, userIdAsFilter, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	period := c.DefaultQuery("size", "all") // all | year | month | week

	var matchFilter bson.D
	switch period {
	case "year":
		matchFilter = bson.D{{"publishedat", bson.D{{"$gte", time.Now().AddDate(-1, 0, 0)}}}}
	case "month":
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

	opts := options.Aggregate().SetBatchSize(100)

	cur, err := handler.collection.Aggregate(handler.ctx, pipeline, opts)
	if err != nil {
		handleDBError(c, err, "error fetching average weights")
		return
	}
	defer cur.Close(handler.ctx)

	var workouts []models.AverageWeight
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
