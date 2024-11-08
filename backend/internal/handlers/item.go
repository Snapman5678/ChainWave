package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

// AddItemHandler handles adding a new item
func AddItemHandler(db *sql.DB, c *gin.Context) {
	var item models.Item

	// Bind the multipart form data to the item struct
	if err := c.ShouldBind(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get roles and role types from the context
	roles, rolesExists := c.Get("roles")
	roleTypes, roleTypesExists := c.Get("roleTypes")
	if !rolesExists || !roleTypesExists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	rolesSlice := roles.([]string)
	roleTypesSlice := roleTypes.([]string)
	if len(rolesSlice) != len(roleTypesSlice) || len(rolesSlice) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Ensure the business admin ID exists in the roles
	var businessAdminId uuid.UUID
	for i, roleType := range roleTypesSlice {
		if roleType == "business_admin" {
			businessAdminId = uuid.MustParse(rolesSlice[i])
			break
		}
	}

	if businessAdminId == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	item.BusinessAdminId = businessAdminId

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	// Save the image to the static/images directory
	imagePath := filepath.Join("static/images", file.Filename)
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}
	//log.Println("Image saved to: ", imagePath

	log.Print("Image saved to: ", imagePath)

	// Set the image URL in the item
	item.ImageURL = "/images/" + file.Filename

	id, err := repository.AddItem(db, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	item.Id = id
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
