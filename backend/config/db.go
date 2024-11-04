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
