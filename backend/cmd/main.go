package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"chainwave/backend/internal/handlers"
	"chainwave/backend/config"
	"chainwave/backend/internal/middleware"
)

func main() {
	// Initialize the database
	log.Println("DATABASE_URL: ", os.Getenv("DATABASE_URL"))
	db, err := config.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a Gin router
	router := gin.Default()

	// Middleware for CORS and JSON Content-Type
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JSONContentTypeMiddleware())

	// User registration and login routes
	router.POST("/api/user/register", func(c *gin.Context) { handlers.RegisterUser(db, c) })
	router.POST("/api/user/login", func(c *gin.Context) { handlers.LoginUser(db, c) })

	// Start the server
	log.Fatal(router.Run(":8000"))
}
