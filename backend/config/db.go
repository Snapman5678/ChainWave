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

	// Create the customer table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		customer_name TEXT,
		contact_info TEXT,
		location TEXT,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		return nil, err
	}

	// Create the business_admin table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS business_admins (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		company_name TEXT,
		contact_info TEXT,
		location TEXT,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		return nil, err
	}

	// Create the transporter table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transporters (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		driver_name TEXT,
		vehicle_details TEXT,
		contact_info TEXT,
		location TEXT,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		return nil, err
	}

	// Create the supplier table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS suppliers (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		supplier_name TEXT,
		contact_info TEXT,
		address TEXT,
		location TEXT,
		user_id UUID,
		FOREIGN KEY (user_id) REFERENCES users(id)
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
