package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	models "github.com/lkhtk/workout-log/models"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/idtoken"
)

var (
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	env            = os.Getenv("STAND")
)

type TokenRequest struct {
	Token string `json:"token"`
}
type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func verifyGoogleToken(token string) (map[string]interface{}, error) {
	payload, err := idtoken.Validate(context.Background(), token, GoogleClientID)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return payload.Claims, nil
}

func (handler *AuthHandler) GoogleAuthHandler(c *gin.Context) {
	var token TokenRequest
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload, err := verifyGoogleToken(token.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	var userId string
	var existingUser models.User
	err = handler.collection.FindOne(context.TODO(), bson.M{"email": payload["email"].(string)}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		user := models.User{
			Name:      payload["name"].(string),
			Email:     payload["email"].(string),
			Picture:   payload["picture"].(string),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		res, err := handler.collection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userId = res.InsertedID.(primitive.ObjectID).String()
	} else {
		userId = existingUser.ID.Hex()
	}
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
	})
	sessionToken := xid.New().String()
	session.Set("token", sessionToken)
	session.Set("email", payload["email"].(string))
	session.Set("user_id", userId)
	if err = session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully", "User": payload})
}

func (handler *AuthHandler) DestroyUserSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "you have successfully logged out"})
	return
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		email := session.Get("email")
		if env == "dev" {
			c.Next()
		} else if sessionToken == nil || email == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
		}
		c.Next()
	}
}

func getCurrentUser(c *gin.Context) (primitive.ObjectID, bson.M, error) {
	user_id := ""
	ok := false
	if env == "dev" {
		user_id = os.Getenv("TEST_USER_ID")
		ok = true
	} else {
		user_id, ok = sessions.Default(c).Get("user_id").(string)
	}
	if !ok || user_id == "" {
		return primitive.ObjectID{}, bson.M{}, errors.New("user is not authenticated")
	}
	objectId, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return primitive.ObjectID{}, bson.M{}, err
	}

	userAsFilter := bson.M{"user_id": objectId}
	userID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return primitive.ObjectID{}, bson.M{}, err
	}
	return userID, userAsFilter, nil
}

func (handler *AuthHandler) DeleteCurrentUser(c *gin.Context) error {
	_, userID, err := getCurrentUser(c)
	if err != nil {
		return err
	}
	session := sessions.Default(c)
	_, err = handler.collection.DeleteOne(c, bson.M{"_id": userID})
	if err != nil {
		return err
	}
	session.Clear()
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
	return nil
}
