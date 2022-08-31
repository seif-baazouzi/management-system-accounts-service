package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUsername struct {
	NewUsername string `json:"newUsername"`
	Password    string `json:"password"`
}