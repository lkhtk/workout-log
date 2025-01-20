package handlers

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	filter := bson.M{
		"measurement_date": measurement.MeasurementDate,
		"user_id":          measurement.UserID,
	}
	update := bson.M{
		"$set": bson.M{
			"body_fat":    measurement.BodyFat,
			"body_weight": measurement.BodyWeight,
			"Neck":        measurement.Neck,
			"Chest":       measurement.Chest,
			"Waist":       measurement.Waist,
			"Hips":        measurement.Hips,
			"UpperArm":    measurement.UpperArm,
			"Forearm":     measurement.Forearm,
			"Thighs":      measurement.Thighs,
			"Calves":      measurement.Calves,
		},
	}
	options := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var updatedMeasurement models.Measurement
	err = handler.collection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&updatedMeasurement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating or inserting measurement"})
		return
	}
	c.JSON(http.StatusOK, updatedMeasurement)
}

func (handler *MongoConnectionHandler) GetAllMeasurements(c *gin.Context) {
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
		SetSort(bson.D{{"measurement_date", -1}})

	cur, err := handler.collection.Find(handler.ctx, filter, findOptions)
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
		"last_page": math.Ceil(float64(total) / float64(perPage)),
	})
}

func validateAndPrepareMeasurement(c *gin.Context) (models.Measurement, error) {
	var measurement models.Measurement
	if err := c.ShouldBindJSON(&measurement); err != nil {
		return measurement, errors.New("invalid measurement data")
	}
	measurement.ID = primitive.NewObjectID()
	userID, err := GetCurrentUser(c)
	if err != nil {
		return measurement, err
	}
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return measurement, err
	}
	measurement.UserID = uid
	return measurement, nil
}
func (handler *MongoConnectionHandler) GetLatestMeasurement(c *gin.Context) {
	userID, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	filter := getFilter(c, userID)

	findOptions := options.Find().
		SetLimit(1).
		SetSort(bson.D{{"measurement_date", -1}})
	cur, err := handler.collection.Find(handler.ctx, filter, findOptions)
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

	c.JSON(http.StatusOK, measurements)
}
