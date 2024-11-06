package models

import "github.com/google/uuid"

// Vehicle struct
type Vehicle struct {
	ID              uuid.UUID `json:"id"`
	TransporterID   uuid.UUID `json:"transporter_id"`
	Make            string    `json:"make"`
	Model           string    `json:"model"`
	Year            int       `json:"year"`
	Location        Location  `json:"location"`
	MaxDistance     float64   `json:"max_distance"`
	MaxCapacity     float64   `json:"max_capacity"`
	CurrentCapacity float64   `json:"current_capacity"`
}

// Location struct
type Location struct {
	ID         uuid.UUID `json:"id"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
	PostalCode string    `json:"postal_code"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
}
