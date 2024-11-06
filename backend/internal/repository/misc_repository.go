package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
	"github.com/google/uuid"
)

// AddLocation adds a new location to the database
func AddLocation(db *sql.DB, location models.Location) error {
	_, err := db.Exec(`INSERT INTO locations (id, address, city, state, country, postal_code, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		location.ID, location.Address, location.City, location.State, location.Country, location.PostalCode, location.Latitude, location.Longitude)
	return err
}

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


// AddVehicle adds a new vehicle to the database
func AddVehicle(db *sql.DB, vehicle models.Vehicle) error {
	_, err := db.Exec(`INSERT INTO vehicles (id, transporter_id, make, model, year, location_id, max_distance, max_capacity, current_capacity) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		vehicle.ID, vehicle.TransporterID, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.Location.ID, vehicle.MaxDistance, vehicle.MaxCapacity, vehicle.CurrentCapacity)
	return err
}

// GetVehicleByID retrieves a vehicle by its ID
func GetVehicleByID(db *sql.DB, id uuid.UUID) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	query := `SELECT id, transporter_id, make, model, year, location_id, max_distance, max_capacity, current_capacity FROM vehicles WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&vehicle.ID, &vehicle.TransporterID, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.Location.ID, &vehicle.MaxDistance, &vehicle.MaxCapacity, &vehicle.CurrentCapacity)
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}
