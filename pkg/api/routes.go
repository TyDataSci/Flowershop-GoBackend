package api

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	//initialize mux router
	ServeRoutes(router)
	return router
}

func ServeRoutes(router *mux.Router) {
	itemRoutes(router)
	userRoutes(router)
	orderRoutes(router)
	orderItemsRoutes(router)
}

func itemRoutes(router *mux.Router) {
	//Routes for Items
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items", CreateItem).Methods("POST")
	router.HandleFunc("/items", UpdateItem).Methods("PUT")
	//router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	//router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
}

func userRoutes(router *mux.Router) {
	//Routes for Users
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{username}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	//router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	//router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
}

func orderRoutes(router *mux.Router) {
	//Routes for Orders
	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", GetOrder).Methods("GET")
	//router.HandleFunc("/orders", CreateOrder).Methods("POST")
	//router.HandleFunc("/orders/{id}", UpdateOrder).Methods("PUT")
	//router.HandleFunc("/orders/{id}", DeleteOrder).Methods("DELETE")
}

func orderItemsRoutes(router *mux.Router) {
	//Routes for Order_Items
	router.HandleFunc("/order_items/{orderid}", GetOrderItems).Methods("GET")
	router.HandleFunc("/order_items", CreateItem).Methods("POST")
	router.HandleFunc("/order_items", UpdateItem).Methods("PUT")
	//router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	//router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
}
