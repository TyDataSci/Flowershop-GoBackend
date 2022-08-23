package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var flowers = []*Item{}
var users = []*User{}
var accounts = []*Account{}
var currentUser *User
var accountMap = make(map[string]*Account)
var userOrderMap = make(map[string][]*Order)

func RemoveItem(slice []*Item, index int) []*Item {
	modified := make([]*Item, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func RemoveUser(slice []*User, index int) []*User {
	modified := make([]*User, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func RemoveAccount(slice []*Account, index int) []*Account {
	modified := make([]*Account, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func RemoveOrder(slice []*Order, index int) []*Order {
	modified := make([]*Order, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func getFlowers(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(flowers)
}

func getUsers(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(users)
}

func getAccounts(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(accounts)
}

func getOrders(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(orders)
}

func getFlower(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for _, item := range flowers {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for _, user := range users {
		if user.Username == params["username"] {
			json.NewEncoder(writer).Encode(user)
			currentUser = user
			return
		}
	}
}

func getAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	//accountID := params["id"]
	for _, account := range accounts {
		if account.ID == params["id"] {
			ord, exists := userOrderMap[currentUser.ID]
			if exists {
				account.Orders = ord
			}
			{
				fmt.Printf("Account %v does not have any orders.\n", currentUser.Username)
			}
			json.NewEncoder(writer).Encode(account)
			return
		}
	}

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

func createFlower(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var flower *Item
	_ = json.NewDecoder(router.Body).Decode(&flower)
	flower.ID = strconv.Itoa(rand.Intn(100000))
	flowers = append(flowers, flower)
	json.NewEncoder(writer).Encode(flowers)
}

func createUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var user *User
	_ = json.NewDecoder(router.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000))
	users = append(users, user)
	json.NewEncoder(writer).Encode(users)
}

func createAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var account *Account
	_ = json.NewDecoder(router.Body).Decode(&account)
	account.ID = currentUser.ID
	account.User = currentUser
	for _, _account := range accounts {
		if _account.ID == currentUser.ID {
			return
		}
	}
	accountMap[account.ID] = account
	accounts = append(accounts, account)
	json.NewEncoder(writer).Encode(accounts)
}

func createOrder(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var order *Order
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

func updateFlower(writer http.ResponseWriter, router *http.Request) {
	//set json content type
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	//json params
	params := mux.Vars(router)
	//loop thru movies,range
	for index, item := range flowers {
		if item.ID == params["id"] {
			flowers = RemoveItem(flowers, index)
			//flowers = append(flowers[:index], flowers[index+1:]...)
			var newFlower *Item
			_ = json.NewDecoder(router.Body).Decode(&newFlower)
			newFlower.ID = params["id"]
			flowers = append(flowers, newFlower)
			json.NewEncoder(writer).Encode(newFlower)
		}
	}
}

func updateUser(writer http.ResponseWriter, router *http.Request) {
	//Remove the previous user and append the modified user with changes
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, user := range users {
		if user.ID == params["id"] {
			users = RemoveUser(users, index)
			var modifyUser *User
			_ = json.NewDecoder(router.Body).Decode(&modifyUser)
			modifyUser.ID = params["id"]
			users = append(users, modifyUser)
			json.NewEncoder(writer).Encode(modifyUser)
		}
	}

}

func updateAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, account := range accounts {
		if account.ID == params["id"] {
			accounts = RemoveAccount(accounts, index)
			var modifyAccount *Account
			_ = json.NewDecoder(router.Body).Decode(&modifyAccount)
			modifyAccount.ID = params["id"]
			accounts = append(accounts, modifyAccount)
			accountMap[modifyAccount.ID] = modifyAccount
			json.NewEncoder(writer).Encode(modifyAccount)
		}
	}
}

func updateOrder(writer http.ResponseWriter, router *http.Request) {
	orders := userOrderMap[currentUser.ID]
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, order := range orders {
		if order.ID == params["id"] {
			orders = RemoveOrder(orders, index)
			var modifyOrder *Order
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

func deleteFlower(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, item := range flowers {
		if item.ID == params["id"] {
			flowers = RemoveItem(flowers, index)
			fmt.Printf("Deleted %v\n", item.ID)
			break
		}
	}
	json.NewEncoder(writer).Encode(flowers)
}

func deleteUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, user := range users {
		if user.ID == params["id"] {
			users = RemoveUser(users, index)
			fmt.Printf("Delete %v\n", user.ID)
			break
		}
	}
}

func deleteAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, account := range accounts {
		if account.ID == params["id"] {
			accounts = RemoveAccount(accounts, index)
			accountMap[account.ID] = nil
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
