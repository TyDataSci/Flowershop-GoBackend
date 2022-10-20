package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/db"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db.Connect()

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", api.Router()))

}
