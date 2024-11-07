package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
	"github.com/google/uuid"
)

// AddItem adds a new item to the database
func AddItem(db *sql.DB, item models.Item) error {
	_, err := db.Exec(`INSERT INTO items (id, name, description, price, weight, dimensions, category, quantity, image_url) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8)`,
		item.Name, item.Description, item.Price, item.Weight, item.Dimensions, item.Category, item.Quantity, item.ImageURL)
	return err
}

// EditItem updates an existing item in the database
func EditItem(db *sql.DB, item models.Item) error {
	_, err := db.Exec(`UPDATE items SET name = $1, description = $2, price = $3, weight = $4, dimensions = $5, category = $6, quantity = $7, image_url = $8 WHERE id = $9`,
		item.Name, item.Description, item.Price, item.Weight, item.Dimensions, item.Category, item.Quantity, item.ImageURL, item.Id)
	return err
}

// GetItemById fetches an item by its ID
func GetItemById(db *sql.DB, itemId uuid.UUID) (models.Item, error) {
	var item models.Item
	err := db.QueryRow(`SELECT id, name, description, price, weight, dimensions, category, quantity, image_url FROM items WHERE id = $1`, itemId).Scan(
		&item.Id, &item.Name, &item.Description, &item.Price, &item.Weight, &item.Dimensions, &item.Category, &item.Quantity, &item.ImageURL)
	if err != nil {
		return item, err
	}
	return item, nil
}

// DeleteItem deletes an item from the database
func DeleteItem(db *sql.DB, itemId uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM items WHERE id = $1`, itemId)
	return err
}
