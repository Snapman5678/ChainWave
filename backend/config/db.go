package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// InitDB initializes the database connection and creates the users table if it doesn't exist.
func InitDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Ensure the UUID extension is available
	if err := ensureUUIDExtension(db); err != nil {
		return nil, err
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT
	)`)
	if err != nil {
		return nil, err
	}

	// Create the locations table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS locations (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		address TEXT,
		city TEXT,
		state TEXT,
		country TEXT,
		postal_code TEXT,
		latitude FLOAT8,
		longitude FLOAT8
	)`)
	if err != nil {
		return nil, err
	}

	// Create the vehicles table without foreign key constraints
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS vehicles (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		transporter_id UUID,
		make TEXT,
		model TEXT,
		year INT,
		latitude FLOAT8,
		longitude FLOAT8,
		max_distance FLOAT8,
		max_capacity FLOAT8,
		current_capacity FLOAT8
	)`)
	if err != nil {
		return nil, err
	}

	// Create the transporters table without foreign key constraints
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transporters (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		driver_name TEXT,
		vehicle_id UUID,
		contact_info TEXT,
		location_id UUID,
		user_id UUID
	)`)
	if err != nil {
		return nil, err
	}

	// Alter the vehicles table to add foreign key constraints
	_, err = db.Exec(`ALTER TABLE vehicles
		ADD CONSTRAINT fk_transporter
		FOREIGN KEY (transporter_id) REFERENCES transporters(id)
	`)
	if err != nil {
		return nil, err
	}

	// Alter the transporters table to add foreign key constraints
	_, err = db.Exec(`ALTER TABLE transporters
		ADD CONSTRAINT fk_user
		FOREIGN KEY (user_id) REFERENCES users(id),
		ADD CONSTRAINT fk_vehicle
		FOREIGN KEY (vehicle_id) REFERENCES vehicles(id),
		ADD CONSTRAINT fk_location
		FOREIGN KEY (location_id) REFERENCES locations(id)
	`)
	if err != nil {
		return nil, err
	}

	// Create the customers table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		customer_name TEXT,
		contact_info TEXT,
		location_id UUID,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (location_id) REFERENCES locations(id)
	)`)
	if err != nil {
		return nil, err
	}

	// Create the business_admins table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS business_admins (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		company_name TEXT,
		contact_info TEXT,
		location_id UUID,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (location_id) REFERENCES locations(id)
	)`)
	if err != nil {
		return nil, err
	}

	// Create the suppliers table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS suppliers (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		supplier_name TEXT,
		contact_info TEXT,
		address TEXT,
		location_id UUID,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (location_id) REFERENCES locations(id)
	)`)
	if err != nil {
		return nil, err
	}

	// Create the user_roles table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user_roles (
		user_id UUID NOT NULL,
		customer_id UUID,
		business_admin_id UUID,
		transporter_id UUID,
		supplier_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (customer_id) REFERENCES customers(id),
		FOREIGN KEY (business_admin_id) REFERENCES business_admins(id),
		FOREIGN KEY (transporter_id) REFERENCES transporters(id),
		FOREIGN KEY (supplier_id) REFERENCES suppliers(id)
	)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ensureUUIDExtension ensures that the uuid-ossp extension is enabled in the database.
func ensureUUIDExtension(db *sql.DB) error {
	// Check if the extension is already available
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (
		SELECT 1
		FROM pg_extension
		WHERE extname = 'uuid-ossp'
	);`).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check uuid-ossp extension: %w", err)
	}

	// Create the extension only if it does not exist
	if !exists {
		_, err := db.Exec(`CREATE EXTENSION "uuid-ossp";`)
		if err != nil {
			return fmt.Errorf("failed to create uuid-ossp extension: %w", err)
		}
	}

	return nil
}
