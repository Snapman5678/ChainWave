package handlers

import (
	"chainwave/backend/internal/models"
	"chainwave/backend/internal/repository"
	"database/sql"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/textproto"
	"fmt"
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

// GetItemHandler handles fetching an item by its ID with details
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

	// Prepare multipart writer
	multipartWriter := multipart.NewWriter(c.Writer)
	c.Writer.Header().Set("Content-Type", multipartWriter.FormDataContentType())

	// Add item as JSON part
	jsonPart, err := multipartWriter.CreateFormField("item")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form field"})
		return
	}
	itemJSON, err := json.Marshal(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal item"})
		return
	}
	jsonPart.Write(itemJSON)

	// Add image as a separate part
	imagePath := filepath.Join("static/images", filepath.Base(item.ImageURL))
	file, err := os.Open(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	file.Seek(0, io.SeekStart)

	partHeaders := textproto.MIMEHeader{}
	partHeaders.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, filepath.Base(item.ImageURL)))
	partHeaders.Set("Content-Type", contentType)
	imagePart, err := multipartWriter.CreatePart(partHeaders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create image part"})
		return
	}
	io.Copy(imagePart, file)

	multipartWriter.Close()
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

// GetItemCountHandler handles fetching the total number of items
func GetItemCountHandler(db *sql.DB, c *gin.Context) {
	count, err := repository.GetItemCount(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// GetItemsByCategoryHandler handles fetching items by category and includes images in the multipart response
func GetItemsByCategoryHandler(db *sql.DB, c *gin.Context, category string, limit, offset int) {
	// The handler now receives category as a parameter from the query
	items, err := repository.GetItemsByCategory(db, category, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Prepare multipart writer
	multipartWriter := multipart.NewWriter(c.Writer)
	c.Writer.Header().Set("Content-Type", multipartWriter.FormDataContentType())

	// Add items as JSON part
	jsonPart, err := multipartWriter.CreateFormField("items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form field"})
		return
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal items"})
		return
	}
	_, err = jsonPart.Write(itemsJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write items JSON"})
		return
	}

	// Add each image as a separate part
	for _, item := range items {
		imagePath := filepath.Join("static/images", filepath.Base(item.ImageURL))
		file, err := os.Open(imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
			return
		}
		defer file.Close()

		// Determine MIME type
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
			return
		}
		contentType := http.DetectContentType(buffer)
		file.Seek(0, io.SeekStart) // Reset file pointer

		// Create image part with appropriate Content-Type
		partHeaders := textproto.MIMEHeader{}
		partHeaders.Set("Content-Disposition", fmt.Sprintf(`form-data; name="images"; filename="%s"`, filepath.Base(item.ImageURL)))
		partHeaders.Set("Content-Type", contentType)
		imagePart, err := multipartWriter.CreatePart(partHeaders)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create image part"})
			return
		}
		_, err = io.Copy(imagePart, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy image data"})
			return
		}
	}

	multipartWriter.Close()
}