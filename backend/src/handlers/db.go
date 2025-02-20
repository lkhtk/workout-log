package handlers

import (
	"net/http"

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

	return filter
}
