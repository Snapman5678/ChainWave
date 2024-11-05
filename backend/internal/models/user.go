package models

import "github.com/google/uuid"

// User struct
type User struct {
	Id       uuid.UUID `json:"id"`       // Change type to uuid.UUID
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
