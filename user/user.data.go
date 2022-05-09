package user

import (
	"context"
	"gotraining.com/database"
	"log"
	"time"
)

func GetUser(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT username, password FROM "user" WHERE username = $1`
	row := database.DbConn.QueryRowContext(ctx, query, username)

	if err := row.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	user := &User{}
	err := row.Scan(
		&user.Username,
		&user.Password)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}
