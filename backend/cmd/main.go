package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"chainwave/backend/internal/handlers"
	"chainwave/backend/config"
)

func main() {
	// Initialize the database
	db, err := config.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a Gin router
	router := gin.Default()

	// Middleware for CORS and JSON Content-Type
	router.Use(CORSMiddleware())
	router.Use(JSONContentTypeMiddleware())

	// Define user-related routes
	router.GET("/api/go/users", func(c *gin.Context) { handlers.GetUsers(db, c) })
	router.POST("/api/go/users", func(c *gin.Context) { handlers.CreateUser(db, c) })
	router.GET("/api/go/users/:id", func(c *gin.Context) { handlers.GetUser(db, c) })
	router.PUT("/api/go/users/:id", func(c *gin.Context) { handlers.UpdateUser(db, c) })
	router.DELETE("/api/go/users/:id", func(c *gin.Context) { handlers.DeleteUser(db, c) })

	// Start the server
	log.Fatal(router.Run(":8000"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
			return
		}
		c.Next()
	}
}

func JSONContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
