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

	// Middleware for CORS, JSON Content-Type, and Auth
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JSONContentTypeMiddleware())

	// User registration and login routes
	router.POST("/api/user/register", func(c *gin.Context) { handlers.RegisterUser(db, c) })
	router.POST("/api/user/login", func(c *gin.Context) { handlers.LoginUser(db, c) })


	// Additional routes for customer, business admin, transporter, and supplier
	
	// Authenticated routes (protected by JWT middleware)
	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.AuthMiddleware("your_secret_key")) // Replace with your actual secret key

	// Routes that require JWT authentication
	authRoutes.POST("/customer", func(c *gin.Context) { handlers.AddCustomerHandler(db, c) })
	authRoutes.PUT("/customer/:id", func(c *gin.Context) { handlers.EditCustomerHandler(db, c) })
	authRoutes.POST("/business-admin", func(c *gin.Context) { handlers.AddBusinessAdminHandler(db, c) })
	authRoutes.PUT("/business-admin/:id", func(c *gin.Context) { handlers.EditBusinessAdminHandler(db, c) })
	authRoutes.POST("/transporter", func(c *gin.Context) { handlers.AddTransporterHandler(db, c) })
	authRoutes.PUT("/transporter/:id", func(c *gin.Context) { handlers.EditTransporterHandler(db, c) })
	authRoutes.POST("/supplier", func(c *gin.Context) { handlers.AddSupplierHandler(db, c) })
	authRoutes.PUT("/supplier/:id", func(c *gin.Context) { handlers.EditSupplierHandler(db, c) })

	// Start the server
	log.Fatal(router.Run(":8000"))
}
