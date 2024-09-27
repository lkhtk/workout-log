package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func main() {
	fmt.Println("Hello world")
	server = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "https://lkhtk.me"}
	config.AllowMethods = []string{"GET"}
	config.AllowCredentials = true
	server.Use(cors.New(config))
	router := server.Group("/api")
	router.GET("/health", health)
	server.Run(":8000")
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{"state": "ok"})
}
