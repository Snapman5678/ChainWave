package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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

// Update a user's email in the database
func UpdateEmail(db *sql.DB, userID uuid.UUID, newEmail string) error {
	// Get current email
	var currentEmail string
	query := `SELECT email FROM users WHERE id = $1`
	err := db.QueryRow(query, userID).Scan(&currentEmail)
	if err != nil {
		return err
	}

	// Check if the new email is the same as the current email
	if currentEmail == newEmail {
		return fmt.Errorf("new email is the same as the current email")
	}

	// Check if email already exists
	exists, err := IsEmailExists(db, newEmail)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email already exists")
	}

	query = `UPDATE users SET email = $1 WHERE id = $2`
	_, err = db.Exec(query, newEmail, userID)
	return err
}

// Update a user's username in the database
func UpdateUsername(db *sql.DB, userID uuid.UUID, newUsername string) error {
	// Get current username
	var currentUsername string
	query := `SELECT username FROM users WHERE id = $1`
	err := db.QueryRow(query, userID).Scan(&currentUsername)
	if err != nil {
		return err
	}

	// Check if the new username is the same as the current username
	if currentUsername == newUsername {
		return fmt.Errorf("new username is the same as the current username")
	}

	// Check if username already exists
	exists, err := IsUsernameExists(db, newUsername)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("username already exists")
	}

	query = `UPDATE users SET username = $1 WHERE id = $2`
	_, err = db.Exec(query, newUsername, userID)
	return err
}

// Update a user's password in the database
func UpdatePassword(db *sql.DB, userID uuid.UUID, newPassword string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := db.Exec(query, newPassword, userID)
	return err
}