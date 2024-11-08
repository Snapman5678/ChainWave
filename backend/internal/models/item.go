package models

import "github.com/google/uuid"

// Item struct
type Item struct {
	Id              uuid.UUID `form:"id"`
	BusinessAdminId uuid.UUID `form:"business_admin_id"`
	Name            string    `form:"name"`
	Description     string    `form:"description"`
	Price           float64   `form:"price"`
	Weight          float64   `form:"weight"`
	Dimensions      string    `form:"dimensions"`
	Category        string    `form:"category"`
	Quantity        int       `form:"quantity"`
	ImageURL        string    `form:"image_url"`
}
