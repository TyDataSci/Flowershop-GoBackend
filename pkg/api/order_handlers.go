package api

import (
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var userOrderMap = make(map[string][]*models.Order)

func GetUserOrderMap() map[string][]*models.Order {
	return userOrderMap
}

func SetUserOrderMap(_userOrderMap map[string][]*models.Order) {
	userOrderMap = _userOrderMap
}

func RemoveOrder(slice []*models.Order, index int) []*models.Order {
	modified := make([]*models.Order, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func getOrders(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(orders)
}

func getOrder(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for _, order := range orders {
		if order.ID == params["id"] {
			json.NewEncoder(writer).Encode(order)
			return
		}
	}
}

func createOrder(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var order *models.Order
	_ = json.NewDecoder(router.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(100000))
	order.UserID = currentUser.ID
	orders = append(orders, order)
	userOrderMap[currentUser.ID] = orders
	for index, account := range accounts {
		if account.ID == currentUser.ID {
			modifyAccount := accountMap[currentUser.ID]
			modifyAccount.ID = currentUser.ID
			modifyAccount.Orders = orders
			accounts = RemoveAccount(accounts, index)
			accounts = append(accounts, modifyAccount)
			accountMap[modifyAccount.ID] = modifyAccount
		}
	}
	json.NewEncoder(writer).Encode(orders)
}

func updateOrder(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, order := range orders {
		if order.ID == params["id"] {
			orders = RemoveOrder(orders, index)
			var modifyOrder *models.Order
			_ = json.NewDecoder(router.Body).Decode(&modifyOrder)
			modifyOrder.ID = params["id"]
			orders = append(orders, modifyOrder)
			userOrderMap[currentUser.ID] = orders
			for index, account := range accounts {
				if account.ID == currentUser.ID {
					modifyAccount := accountMap[currentUser.ID]
					modifyAccount.ID = currentUser.ID
					modifyAccount.Orders = orders
					accounts = RemoveAccount(accounts, index)
					accounts = append(accounts, modifyAccount)
					accountMap[modifyAccount.ID] = modifyAccount
				}
			}
			json.NewEncoder(writer).Encode(modifyOrder)
		}
	}
}

func deleteOrder(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, order := range orders {
		if order.ID == params["id"] {
			orders = RemoveOrder(orders, index)
			userOrderMap[currentUser.ID] = orders
		}
	}
	for index, account := range accounts {
		if account.ID == currentUser.ID {
			modifyAccount := accountMap[currentUser.ID]
			modifyAccount.ID = currentUser.ID
			modifyAccount.Orders = orders
			accounts = RemoveAccount(accounts, index)
			accounts = append(accounts, modifyAccount)
			accountMap[modifyAccount.ID] = modifyAccount
		}
	}
}
