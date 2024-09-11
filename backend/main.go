package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// main function
func main() {
	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Create a Gin router
	router := gin.Default()

	// Middleware for CORS
	router.Use(CORSMiddleware())
	// Middleware for JSON Content-Type
	router.Use(JSONContentTypeMiddleware())

	// Define routes
	router.GET("/api/go/users", func(c *gin.Context) { getUsers(db, c) })
	router.POST("/api/go/users", func(c *gin.Context) { createUser(db, c) })
	router.GET("/api/go/users/:id", func(c *gin.Context) { getUser(db, c) })
	router.PUT("/api/go/users/:id", func(c *gin.Context) { updateUser(db, c) })
	router.DELETE("/api/go/users/:id", func(c *gin.Context) { deleteUser(db, c) })

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

// getUsers function
func getUsers(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// getUser by id
func getUser(db *sql.DB, c *gin.Context) {
	id := c.Param("id")
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	user := User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// createUser function
func createUser(db *sql.DB, c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	err := db.QueryRow("INSERT INTO users(name,email) VALUES($1,$2) RETURNING id", user.Name, user.Email).Scan(&user.Id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// updateUser function
func updateUser(db *sql.DB, c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	id := c.Param("id")
	_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var updatedUser User
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Email)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// deleteUser function
func deleteUser(db *sql.DB, c *gin.Context) {
	id := c.Param("id")

	var user User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
