package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(db *sql.DB, c *gin.Context) {
	users, err := repository.GetAllUsers(db)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(db *sql.DB, c *gin.Context) {
	id := c.Param("id")
	user, err := repository.GetUserByID(db, id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(db *sql.DB, c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if err := repository.CreateUser(db, &user); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(db *sql.DB, c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	id := c.Param("id")
	if err := repository.UpdateUser(db, &user, id); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(db *sql.DB, c *gin.Context) {
	id := c.Param("id")
	if err := repository.DeleteUser(db, id); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
