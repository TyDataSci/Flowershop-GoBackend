package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetItems(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	items, _ := db.GetItems()
	json.NewEncoder(writer).Encode(items)
}

func CreateItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	db.CreateItem(params["type"], params["description"], params["price"], params["image"])
}

func UpdateItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var modifyItem models.Item
	json.NewDecoder(router.Body).Decode(&modifyItem)
	db.UpdateItem(modifyItem)
}

func GetItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "https://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	params := mux.Vars(router)
	itemID, _ := strconv.Atoi(params["id"])
	item, _ := db.GetItem(itemID)
	json.NewEncoder(writer).Encode(item)
}

func DeleteItem(writer http.ResponseWriter, router *http.Request) {
}
