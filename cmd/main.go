package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/middleware"
	"Flowershop-GoBackend/pkg/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db.Connect()
	var user models.User
	user.Password = "Password"
	bytes, _ := bcrypt.GenerateFromPassword([]byte("Password"), 14)
	err := middleware.HashwordCompare(string(bytes), user)
	if err == nil {
		fmt.Println("Success compare")
	} else {
		fmt.Printf("Fail compare %v", err)
	}
	fmt.Println(string(bytes))
	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", api.Router()))

}
