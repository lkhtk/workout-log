package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
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
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		return
	}
	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("token", sessionToken)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully", "User": payload})
}

func (handler *AuthHandler) DestroyUserSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "you have successfully logged out"})
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		log.Panic(sessionToken)
		if sessionToken == nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Unauthorized"})
			c.Abort()
		}
		c.Next()
	}
}
