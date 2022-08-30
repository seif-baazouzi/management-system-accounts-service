package db

import (
	"fmt"
	"os"
)

func Migrations() {
	conn := GetPool()
	defer ClosePool(conn)

	_, err := conn.Exec(
		`CREATE TABLE IF NOT EXISTS users (
			userID UUID PRIMARY KEY,
			username VARCHAR NOT NULL, 
			password VARCHAR NOT NULL,
			createdAt TIMESTAMP DEFAULT NOW() 
		)`,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
