package dev

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/models"
)

func InitializeMockData() {

	var currentUser = api.GetCurrentUser()
	//Mock data for Flowers
	var flowers []*models.Item
	flowers = api.GetFlowersArray()
	flowers = append(flowers, &models.Item{ID: "1", ItemType: "Flower", Description: "Red Rose", Price: 10.00, Inventory: &models.Inventory{Count: 0, MinimumCount: 0, Supplier: "Reynolds", LeadDays: 1, Cost: 4.50}})
	flowers = append(flowers, &models.Item{ID: "2", ItemType: "Flower", Description: "White Rose", Price: 10.00, Inventory: &models.Inventory{Count: 0, MinimumCount: 0, Supplier: "Reynolds", LeadDays: 1, Cost: 4.50}})
	api.SetFlowersArray(flowers)

	//Mock data for Users
	var users []*models.User
	users = api.GetUsersArray()
	users = append(users, currentUser)
	users = append(users, &models.User{ID: "1", Username: "tssand", Password: "Password"})
	users = append(users, &models.User{ID: "2", Username: "wit23", Password: "Password2"})
	users = append(users, &models.User{ID: "3", Username: "princessisla", Password: "Password3"})
	api.SetUsersArray(users)

	//Mock data for Accounts
	var accounts []*models.Account
	accounts = api.GetAccountsArray()
	accounts = append(accounts, &models.Account{ID: "1", User: &models.User{ID: "1", Username: "tssand", Password: "Password"}, FirstName: "Ty", LastName: "Sanders", Email: "tssanders2@gmail.com", Phone: "555-555-5555", PaymentMethod: "Debit", Orders: []*models.Order{}})
	api.SetAccountsArray(accounts)

	var accountMap = api.GetAccountMap()
	accountMap["1"] = &models.Account{ID: "1", User: &models.User{ID: "1", Username: "tssand", Password: "Password"}, FirstName: "Ty", LastName: "Sanders", Email: "tssanders2@gmail.com", Phone: "555-555-5555", PaymentMethod: "Debit", Orders: []*models.Order{}}
	api.SetAccountMap(accountMap)

	//Mock data for Orders
	var ord = make([]*models.Order, 0)
	ord = append(ord, &models.Order{ID: "1", Date: "07-04-2022", UserID: "1", OrderItems: "Red Rose", DeliveryType: "Deliver", Note: "", Instructions: "", TotalCost: 10.00})
	var userOrderMap = api.GetUserOrderMap()
	userOrderMap["1"] = ord
	api.SetUserOrderMap(userOrderMap)
}
