package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/dev"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//initialize router to handle api calls
	dev.InitializeMockData()
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", api.Router()))
}
