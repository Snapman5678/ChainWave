package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
	"github.com/google/uuid"
)

// AddItem adds a new item to the database
func AddItem(db *sql.DB, item models.Item) (uuid.UUID, error) {
	var id uuid.UUID
	err := db.QueryRow(`INSERT INTO items (id, business_admin_id, name, description, price, weight, dimensions, category, quantity, image_url) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		item.BusinessAdminId, item.Name, item.Description, item.Price, item.Weight, item.Dimensions, item.Category, item.Quantity, item.ImageURL).Scan(&id)
	return id, err
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

// GetItemCount fetches the total number of items in the database
func GetItemCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM items`).Scan(&count)
	return count, err
}

// GetItemsByCategory fetches a list of items from the database
func GetItemsByCategory(db *sql.DB, category string, offset int, limit int) ([]models.Item, error) {
	var rows *sql.Rows
	var err error
	if category == "" {
		rows, err = db.Query(`SELECT id, name, description, price, weight, dimensions, category, quantity, image_url FROM items OFFSET $1 LIMIT $2`, offset, limit)
	} else {
		rows, err = db.Query(`SELECT id, name, description, price, weight, dimensions, category, quantity, image_url FROM items WHERE category = $1 OFFSET $2 LIMIT $3`, category, offset, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.Weight, &item.Dimensions, &item.Category, &item.Quantity, &item.ImageURL); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}