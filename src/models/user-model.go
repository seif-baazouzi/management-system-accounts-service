package models

import (
	"accounts-service/src/db"
	"accounts-service/src/utils"

	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsersByUsername(username string) ([]User, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var usersList []User

	rows, err := conn.Query("SELECT userID, username, password FROM users WHERE username = $1", username)

	if err != nil {
		return usersList, err
	}

	defer rows.Close()

	if rows.Next() {
		var user User
		rows.Scan(&user.UserID, &user.Username, &user.Password)
		usersList = append(usersList, user)
	}

	return usersList, nil
}

func CreateUser(user *User) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	userID := uuid.New()
	hash := utils.Hash(user.Password)

	_, err := conn.Exec(
		"INSERT INTO users VALUES($1, $2, $3)",
		userID,
		user.Username,
		hash,
	)

	if err != nil {
		return err
	}

	return nil
}
