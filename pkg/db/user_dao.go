package db

import (
	"Flowershop-GoBackend/pkg/models"
	"context"
	"fmt"
)

func GetUser(paramUsername string) (models.User, error) {
	var user models.User

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return user, err
	}
	row := db().QueryRowContext(context.Background(), "SELECT * FROM users WHERE username = $1",
		paramUsername)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return user, err
	}

	if err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Password); err != nil {
		fmt.Println("row.Scan", err)
		return user, err
	}

	fmt.Printf("userid: %v, username: %v, name: %v, password: %v\n", user.ID, user.Username, user.Name, user.Password)
	return user, nil
}

func CreateUser(paramUsername string, paramName string, paramPassword string) (models.User, error) {
	var user models.User
	err := db().QueryRowContext(context.Background(),
		"INSERT INTO users(username,name,password) VALUES($1,$2,$3)RETURNING *",
		paramUsername,
		paramName,
		paramPassword).Scan(&user.ID, &user.Username, &user.Name, &user.Password)
	if err != nil {
		fmt.Println("db.QueryRowContext", err)
		return user, err
	}
	fmt.Println(user.ID, user.Username, user.Name, user.Password)
	return user, nil
}
