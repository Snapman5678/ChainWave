package repository

import (
	"chainwave/backend/internal/models"
	"database/sql"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(db *sql.DB, id string) (models.User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	var user models.User
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, user *models.User) error {
	return db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.Id)
}

func UpdateUser(db *sql.DB, user *models.User, id string) error {
	_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, id)
	return err
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
