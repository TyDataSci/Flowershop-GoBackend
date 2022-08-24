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
	router.HandleFunc("/flowers", getFlowers).Methods("GET")
	router.HandleFunc("/flowers/{id}", getFlower).Methods("GET")
	router.HandleFunc("/flowers", createFlower).Methods("POST")
	router.HandleFunc("/flowers/{id}", updateFlower).Methods("PUT")
	router.HandleFunc("/flowers/{id}", deleteFlower).Methods("DELETE")
}

func userRoutes(router *mux.Router) {
	//Routes for Users
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{username}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
}

func accountRoutes(router *mux.Router) {
	//Routes for Accounts
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", getAccount).Methods("GET")
	router.HandleFunc("/accounts", createAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", updateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", deleteAccount).Methods("DELETE")
}

func orderRoutes(router *mux.Router) {
	//Routes for Orders
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")
}
