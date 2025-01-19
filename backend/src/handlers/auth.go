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
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/idtoken"
)

var (
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
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
		userId = res.InsertedID.(string)
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
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		email := session.Get("email")
		if sessionToken == nil || email == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
		}
		c.Next()
	}
}

func GetCurrentUser(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	user_id, ok := session.Get("user_id").(string)
	if !ok || user_id == "" {
		return "", errors.New("user is not authenticated")
	}
	return user_id, nil
}
