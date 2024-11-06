package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

// Registration handler
func RegisterUser(db *sql.DB, c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Check if email already exists
	exists, err := repository.IsEmailExists(db, user.Email)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash the password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Print(err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	// 	return
	// }
	// user.Password = string(hashedPassword)

	if err := repository.CreateUser(db, &user); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": user.Id})
}

// Login handler
func LoginUser(db *sql.DB, c *gin.Context) {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    
    // Bind JSON input to the loginData struct
    if err := c.BindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
        return
    }

    // Retrieve user by email
    user, err := repository.GetUserByEmail(db, loginData.Email)
    if err != nil {
        // If the user does not exist, return unauthorized error
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Since you have removed hashing, you can directly check if the password matches
    if user.Password != loginData.Password { // Replace with actual password check logic
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // If login is successful
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.Id,
    })
    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString, "user_id": user.Id})
}
