package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddItemHandler handles adding a new item
func AddItemHandler(db *sql.DB, c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.AddItem(db, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// EditItemHandler handles editing an existing item
func EditItemHandler(db *sql.DB, c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.EditItem(db, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// GetItemHandler handles fetching an item by its ID
func GetItemHandler(db *sql.DB, c *gin.Context) {
	itemId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}
	item, err := repository.GetItemById(db, itemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// DeleteItemHandler handles deleting an item
func DeleteItemHandler(db *sql.DB, c *gin.Context) {
	itemId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}
	if err := repository.DeleteItem(db, itemId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
