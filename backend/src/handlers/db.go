package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleDBError(c *gin.Context, err error, notFoundMessage string) {
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getFilter(c *gin.Context, userId string) bson.M {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil
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
