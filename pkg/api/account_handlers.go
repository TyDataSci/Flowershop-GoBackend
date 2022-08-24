package api

import (
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var accounts = []*models.Account{}
var accountMap = make(map[string]*models.Account)

func GetAccountsArray() []*models.Account {
	return accounts
}

func SetAccountsArray(_accounts []*models.Account) {
	accounts = _accounts
}

func GetAccountMap() map[string]*models.Account {
	return accountMap
}

func SetAccountMap(_accountMap map[string]*models.Account) {
	accountMap = _accountMap
}

func RemoveAccount(slice []*models.Account, index int) []*models.Account {
	modified := make([]*models.Account, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func getAccounts(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(accounts)
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

func createAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var account *models.Account
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

func updateAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, account := range accounts {
		if account.ID == params["id"] {
			accounts = RemoveAccount(accounts, index)
			var modifyAccount *models.Account
			_ = json.NewDecoder(router.Body).Decode(&modifyAccount)
			modifyAccount.ID = params["id"]
			accounts = append(accounts, modifyAccount)
			accountMap[modifyAccount.ID] = modifyAccount
			json.NewEncoder(writer).Encode(modifyAccount)
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
