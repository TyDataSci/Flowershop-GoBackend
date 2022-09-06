package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"net/http"

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
}

func DeleteItem(writer http.ResponseWriter, router *http.Request) {
}
