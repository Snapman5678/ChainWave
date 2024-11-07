package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
	"github.com/google/uuid"
)


// GetLocationByID retrieves a location by its ID
func GetLocationByID(db *sql.DB, id uuid.UUID) (*models.Location, error) {
	var location models.Location
	query := `SELECT id, address, city, state, country, postal_code, latitude, longitude FROM locations WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&location.ID, &location.Address, &location.City, &location.State, &location.Country, &location.PostalCode, &location.Latitude, &location.Longitude)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// GetVehicleByID retrieves a vehicle by its ID
func GetVehicleByID(db *sql.DB, id uuid.UUID) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	query := `SELECT id, transporter_id, make, model, year, latitude, longitude, max_distance, max_capacity, current_capacity FROM vehicles WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&vehicle.ID, &vehicle.TransporterID, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.Latitude, &vehicle.Longitude, &vehicle.MaxDistance, &vehicle.MaxCapacity, &vehicle.CurrentCapacity)
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}
