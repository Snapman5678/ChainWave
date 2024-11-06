package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
)

// Create a new user in the database
func CreateUser(db *sql.DB, user *models.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	return db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.Id)
}

// Get a user by email (for login)
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Check if an email already exists in the database
func IsEmailExists(db *sql.DB, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

// Check if a username already exists in the database
func IsUsernameExists(db *sql.DB, username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := db.QueryRow(query, username).Scan(&exists)
	return exists, err
}