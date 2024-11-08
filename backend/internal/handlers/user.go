package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"log"
	"net/http"
	"github.com/google/uuid"
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

	// Check if username already exists
	exists, err = repository.IsUsernameExists(db, user.Username)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
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

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
	})
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "User registered successfully",
		"user_id":  user.Id,
		"username": user.Username,
		"email":    user.Email,
		"token":    tokenString,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"token":    tokenString,
		"user_id":  user.Id,
		"username": user.Username,
		"email":    user.Email,
	})
}

// Update email handler
func UpdateEmailHandler(db *sql.DB, c *gin.Context) {
	var emailData struct {
		Email string `json:"email"`
	}
	if err := c.BindJSON(&emailData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Retrieve the user ID from the context
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Update email
	err = repository.UpdateEmail(db, uid, emailData.Email)
	if err != nil {
		log.Print(err)
		if err.Error() == "new email is the same as the current email" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "New email is the same as the current email"})
		} else if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email updated successfully"})
}

// Update username handler
func UpdateUsernameHandler(db *sql.DB, c *gin.Context) {
	var usernameData struct {
		Username string `json:"username"`
	}
	if err := c.BindJSON(&usernameData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Retrieve the user ID from the context
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Update username
	err = repository.UpdateUsername(db, uid, usernameData.Username)
	if err != nil {
		log.Print(err)
		if err.Error() == "new username is the same as the current username" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "New username is the same as the current username"})
		} else if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Username updated successfully"})
}

// Update password handler
func UpdatePasswordHandler(db *sql.DB, c *gin.Context) {
	var passwordData struct {
		Password string `json:"password"`
	}
	if err := c.BindJSON(&passwordData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Retrieve the user ID from the context
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	uid, err := uuid.Parse(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Update password
	err = repository.UpdatePassword(db, uid, passwordData.Password)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
