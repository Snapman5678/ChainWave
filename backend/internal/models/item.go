package models

import "github.com/google/uuid"

// Item struct
type Item struct {
	Id         uuid.UUID `json:"id"`
	BusinessAdminId uuid.UUID `json:"business_admin_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Weight	  float64   `json:"weight"`
	Dimensions string    `json:"dimensions"`
	Category    string    `json:"category"`
	Quantity    int       `json:"quantity"`
	ImageURL    string    `json:"image_url"`
}
