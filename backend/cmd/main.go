package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"chainwave/backend/internal/handlers"
	"chainwave/backend/config"
	"chainwave/backend/internal/middleware"
)

func main() {
	// Initialize the database
	log.Println("DATABASE_URL: ", os.Getenv("DATABASE_URL"))
	db, err := config.InitDB(os.Getenv("DATABASE_URL"))
	if (err != nil) {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a Gin router
	router := gin.Default()

	// Middleware for CORS, JSON Content-Type, and Auth
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JSONContentTypeMiddleware())

	// Serve static files from the "backend/static/images" directory
	router.Static("/images", "./backend/static/images")

	// User registration and login routes
	router.POST("/api/user/register", func(c *gin.Context) { handlers.RegisterUser(db, c) })
	router.POST("/api/user/login", func(c *gin.Context) { handlers.LoginUser(db, c) })


	// Additional routes for customer, business admin, transporter, and supplier
	
	// Authenticated routes (protected by JWT middleware)
	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.AuthMiddleware("your_secret_key")) // Replace with your actual secret key
	authRoutes.Use(middleware.RoleSwitchMiddleware(db))

	// Routes that require JWT authentication
	authRoutes.POST("/customer", func(c *gin.Context) { handlers.AddCustomerHandler(db, c) })
	authRoutes.PUT("/customer/:id", func(c *gin.Context) { handlers.EditCustomerHandler(db, c) })
	authRoutes.POST("/business-admin", func(c *gin.Context) { handlers.AddBusinessAdminHandler(db, c) })
	authRoutes.PUT("/business-admin/:id", func(c *gin.Context) { handlers.EditBusinessAdminHandler(db, c) })
	authRoutes.POST("/transporter", func(c *gin.Context) { handlers.AddTransporterHandler(db, c) })
	authRoutes.PUT("/transporter/:id", func(c *gin.Context) { handlers.EditTransporterHandler(db, c) })
	authRoutes.POST("/supplier", func(c *gin.Context) { handlers.AddSupplierHandler(db, c) })
	authRoutes.PUT("/supplier/:id", func(c *gin.Context) { handlers.EditSupplierHandler(db, c) })
	authRoutes.GET("/role", func(c *gin.Context) { handlers.GetRolesHandler(db, c) })

	// Routes to update email,username and password for a user
	authRoutes.PUT("/user/email", func(c *gin.Context) { handlers.UpdateEmailHandler(db, c) })
	authRoutes.PUT("/user/username", func(c *gin.Context) { handlers.UpdateUsernameHandler(db, c) })
	authRoutes.PUT("/user/password", func(c *gin.Context) { handlers.UpdatePasswordHandler(db, c) })

    // Authenticated routes for roles and puts role ids in the context
	authRoleRoutes := router.Group("/api/roles")
	authRoleRoutes.Use(middleware.AuthAdminMiddleware("your_secret_key", db)) // Replace with your actual secret key

	// Item-related routes
	authRoleRoutes.GET("/item/count", func(c *gin.Context) { handlers.GetItemCountHandler(db, c) })
	itemRoutes := authRoleRoutes.Group("/item")

	// Middleware for form data
	itemRoutes.Use(middleware.FormContentTypeMiddleware())
	itemRoutes.POST("/", func(c *gin.Context) { handlers.AddItemHandler(db, c) })
	itemRoutes.PUT("/:id", func(c *gin.Context) { handlers.EditItemHandler(db, c) }) // Added PUT route for editing

	// Add a new GET route that uses query parameters
	itemRoutes.GET("/", func(c *gin.Context) {
		category := c.Query("category")
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
		handlers.GetItemsByCategoryHandler(db, c, category, limit, offset)
	})

	// Start the server
	log.Fatal(router.Run(":8000"))
}
