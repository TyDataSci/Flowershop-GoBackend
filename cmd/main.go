package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var userOrderMap = make(map[string][]*models.Order)

	//initialize router to handle api calls
	router := mux.NewRouter()
	api.ServeRoutes(router)
	var curUser = api.GetCurrentUser()
	//Mock data for Flowers
	var flowers []*models.Item
	flowers = api.GetFlowers()
	flowers = append(flowers, &models.Item{ID: "1", ItemType: "Flower", Description: "Red Rose", Price: 10.00, Inventory: &models.Inventory{Count: 0, MinimumCount: 0, Supplier: "Reynolds", LeadDays: 1, Cost: 4.50}})
	flowers = append(flowers, &models.Item{ID: "2", ItemType: "Flower", Description: "White Rose", Price: 10.00, Inventory: &models.Inventory{Count: 0, MinimumCount: 0, Supplier: "Reynolds", LeadDays: 1, Cost: 4.50}})

	//Mock data for Users
	var users []*models.User
	users = api.GetUsers()
	users = append(users, curUser)
	users = append(users, &models.User{ID: "1", Username: "tssand", Password: "Password"})
	users = append(users, &models.User{ID: "2", Username: "wit23", Password: "Password2"})
	users = append(users, &models.User{ID: "3", Username: "princessisla", Password: "Password3"})

	//Mock data for Accounts
	var accounts []*models.Account
	accounts = api.GetAccounts()
	accounts = append(accounts, &models.Account{ID: "1", User: &models.User{ID: "1", Username: "tssand", Password: "Password"}, FirstName: "Ty", LastName: "Sanders", Email: "tssanders2@gmail.com", Phone: "555-555-5555", PaymentMethod: "Debit", Orders: []*Order{}})
	accountMap["1"] = &models.Account{ID: "1", User: &models.User{ID: "1", Username: "tssand", Password: "Password"}, FirstName: "Ty", LastName: "Sanders", Email: "tssanders2@gmail.com", Phone: "555-555-5555", PaymentMethod: "Debit", Orders: []*Order{}}

	//Mock data for Orders
	var ord = make([]*models.Order, 0)
	ord = append(ord, &models.Order{ID: "1", Date: "07-04-2022", UserID: "1", OrderItems: "Red Rose", DeliveryType: "Deliver", Note: "", Instructions: "", TotalCost: 10.00})
	userOrderMap["1"] = ord

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
