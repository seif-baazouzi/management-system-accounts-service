package models

import (
	"accounts-service/src/db"
	"accounts-service/src/utils"
	"errors"

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

func GetUserByUserID(userID string) (User, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var user User

	rows, err := conn.Query("SELECT userID, username, password FROM users WHERE userID = $1", userID)

	if err != nil {
		return user, err
	}

	defer rows.Close()

	if !rows.Next() {
		return user, errors.New("User does not exist")
	}

	rows.Scan(&user.UserID, &user.Username, &user.Password)

	return user, nil
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

func CreateUser(user *User) (string, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	userID := uuid.New()
	hash := utils.Hash(user.Password)

	err := conn.QueryRow(
		"INSERT INTO users VALUES($1, $2, $3) RETURNING userID",
		userID,
		user.Username,
		hash,
	).Scan(&user.UserID)

	if err != nil {
		return "", err
	}

	return userID.String(), nil
}

func UpdateUser(user *User) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	hash := utils.Hash(user.Password)

	_, err := conn.Exec(
		"UPDATE users SET username = $1, password = $2 WHERE userID = $3",
		user.Username,
		hash,
		user.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}
