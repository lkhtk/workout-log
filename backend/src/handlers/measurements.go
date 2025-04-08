package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"github.com/lkhtk/workout-log/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const measurementsLimits int64 = 50

func NewMeasurementsHandler(ctx context.Context, collection *mongo.Collection) *MongoConnectionHandler {
	return &MongoConnectionHandler{
		collection: collection,
		ctx:        ctx,
	}
}
func (handler *MongoConnectionHandler) Create(c *gin.Context) {
	measurement, err := validateAndPrepareMeasurement(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.D{
		{"measurementDate", measurement.MeasurementDate},
		{"user_id", measurement.UserID},
	}
	update := bson.D{
		{"$set", bson.D{
			{"bodyFat", measurement.BodyFat},
			{"bodyWeight", measurement.BodyWeight},
			{"Neck", measurement.Neck},
			{"Chest", measurement.Chest},
			{"Waist", measurement.Waist},
			{"Hips", measurement.Hips},
			{"UpperArm", measurement.UpperArm},
			{"Forearm", measurement.Forearm},
			{"Thighs", measurement.Thighs},
			{"Calves", measurement.Calves},
		}},
	}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var updatedMeasurement models.Measurement
	err = handler.collection.FindOneAndUpdate(handler.ctx, filter, update, opts).Decode(&updatedMeasurement)
	if err != nil {
		handleDBError(c, err, "error inserting or updating measurement")
		return
	}
	c.JSON(http.StatusOK, updatedMeasurement)
}

func (handler *MongoConnectionHandler) GetAllMeasurements(c *gin.Context) {
	_, userIDAsAfilter, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	total, err := handler.collection.CountDocuments(handler.ctx, userIDAsAfilter)
	if err != nil {
		handleDBError(c, err, "error counting documents")
		return
	}

	opts := options.Find().
		SetSkip(int64(page-1) * measurementsLimits).
		SetLimit(measurementsLimits).
		SetSort(bson.D{{"measurementDate", -1}}).
		SetProjection(bson.D{{"user_id", 0}, {"id", 0}, {"_id", 0}})

	cur, err := handler.collection.Find(handler.ctx, userIDAsAfilter, opts)
	if err != nil {
		handleDBError(c, err, "error fetching measurements")
		return
	}
	defer cur.Close(handler.ctx)

	var measurements []models.Measurement
	if err := cur.All(handler.ctx, &measurements); err != nil {
		handleDBError(c, err, "error decoding measurements")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      measurements,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total) / float64(measurementsLimits)),
	})
}

func validateAndPrepareMeasurement(c *gin.Context) (models.Measurement, error) {
	var measurement models.Measurement
	if err := c.ShouldBindJSON(&measurement); err != nil {
		return measurement, errors.New("invalid measurement data")
	}
	userIDPrimitive, _, err := getCurrentUser(c)
	if err != nil {
		return measurement, err
	}
	measurement.UserID = userIDPrimitive

	if measurement.BodyWeight == nil || *measurement.BodyWeight <= 0 {
		return measurement, errors.New("invalid BodyWeight data")
	}
	if measurement.BodyFat != nil && (*measurement.BodyFat < 0 || *measurement.BodyFat > 100) {
		return measurement, errors.New("invalid BodyFat data")
	}

	return measurement, nil
}

func (handler *MongoConnectionHandler) GetLatestMeasurement(c *gin.Context) {
	_, userIDAsAfilter, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	opts := options.FindOne().
		SetSort(bson.D{{"measurementDate", -1}}).
		SetProjection(bson.D{{"user_id", 0}, {"id", 0}, {"_id", 0}})

	var measurement models.Measurement
	err = handler.collection.FindOne(handler.ctx, userIDAsAfilter, opts).Decode(&measurement)
	if err != nil {
		handleDBError(c, err, "no measurements found")
		return
	}

	c.JSON(http.StatusOK, measurement)
}

func (handler *MongoConnectionHandler) ExportMeasurements(ctx *gin.Context, zipWriter *utils.ZipWriter) error {
	_, userIDAsAfilter, err := getCurrentUser(ctx)
	if err != nil {
		return err
	}

	opts := options.Find().SetProjection(bson.D{{"user_id", 0}, {"id", 0}, {"_id", 0}})

	cursor, err := handler.collection.Find(ctx, userIDAsAfilter, opts)
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

	return zipWriter.AddFile("measurements.json", data)
}

func (handler *MongoConnectionHandler) CleanupUserMeasurements(c *gin.Context) {
	_, userIDAsAfilter, err := getCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res, err := handler.collection.DeleteMany(handler.ctx, userIDAsAfilter)
	if err != nil {
		handleDBError(c, err, "error deleting user measurements")
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_count": res.DeletedCount})
}
