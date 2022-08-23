package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//initialize router to handle api calls
	router := mux.NewRouter()
	ServeRoutes(router)

	//Mock data for Flowers
	flowers = append(flowers, &Item{ID: "1", ItemType: "Flower", Description: "Red Rose", Price: 10.00, Inventory: &Inventory{Count: 0, MinimumCount: 0, Supplier: "Reynolds", LeadDays: 1, Cost: 4.50}})
	flowers = append(flowers, &Item{ID: "2", ItemType: "Flower", Description: "White Rose", Price: 10.00, Inventory: &Inventory{Count: 0, MinimumCount: 0, Supplier: "Reynolds", LeadDays: 1, Cost: 4.50}})

	//Mock data for Users
	users = append(users, &User{ID: "1", Username: "tssand", Password: "Password"})
	users = append(users, &User{ID: "2", Username: "wit23", Password: "Password2"})
	users = append(users, &User{ID: "3", Username: "princessisla", Password: "Password3"})

	//Mock data for Accounts
	accounts = append(accounts, &Account{ID: "1", User: &User{ID: "1", Username: "tssand", Password: "Password"}, FirstName: "Ty", LastName: "Sanders", Email: "tssanders2@gmail.com", Phone: "555-555-5555", PaymentMethod: "Debit", Orders: []*Order{}})
	accountMap["1"] = &Account{ID: "1", User: &User{ID: "1", Username: "tssand", Password: "Password"}, FirstName: "Ty", LastName: "Sanders", Email: "tssanders2@gmail.com", Phone: "555-555-5555", PaymentMethod: "Debit", Orders: []*Order{}}

	//Mock data for Orders
	var ord = make([]*Order, 0)
	ord = append(ord, &Order{ID: "1", Date: "07-04-2022", UserID: "1", OrderItems: "Red Rose", DeliveryType: "Deliver", Note: "", Instructions: "", TotalCost: 10.00})
	userOrderMap["1"] = ord

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
