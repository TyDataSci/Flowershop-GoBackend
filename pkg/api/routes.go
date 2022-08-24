package api

import (
	"github.com/gorilla/mux"
)

func ServeRoutes(router *mux.Router) {
	flowerRoutes(router)
	userRoutes(router)
	accountRoutes(router)
	orderRoutes(router)
}

func flowerRoutes(router *mux.Router) {
	//Routes for Flowers
	router.HandleFunc("/flowers", GetFlowers).Methods("GET")
	router.HandleFunc("/flowers/{id}", GetFlower).Methods("GET")
	router.HandleFunc("/flowers", CreateFlower).Methods("POST")
	router.HandleFunc("/flowers/{id}", UpdateFlower).Methods("PUT")
	router.HandleFunc("/flowers/{id}", DeleteFlower).Methods("DELETE")
}

func userRoutes(router *mux.Router) {
	//Routes for Users
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{username}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
}

func accountRoutes(router *mux.Router) {
	//Routes for Accounts
	router.HandleFunc("/accounts", GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", GetAccount).Methods("GET")
	router.HandleFunc("/accounts", CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", DeleteAccount).Methods("DELETE")
}

func orderRoutes(router *mux.Router) {
	//Routes for Orders
	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", GetOrder).Methods("GET")
	router.HandleFunc("/orders", CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", DeleteOrder).Methods("DELETE")
}
