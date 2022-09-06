package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func GetOrderItems(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	orderid, _ := strconv.Atoi(params["orderid"])
	order_items, _ := db.GetOrderItems(orderid)
	json.NewEncoder(writer).Encode(order_items)
}

func CreateOrderItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var order_item models.Order_Item
	json.NewDecoder(router.Body).Decode(&order_item)
	db.CreateOrderItem(order_item.OrderID, order_item.ItemID)
	order_items, _ := db.GetOrderItems(order_item.OrderID)
	json.NewEncoder(writer).Encode(order_items)
}

func UpdateOrderItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var modifyOrderItem models.Order_Item
	json.NewDecoder(router.Body).Decode(&modifyOrderItem)
	db.UpdateOrderItem(modifyOrderItem)
	order_items, _ := db.GetOrderItems(modifyOrderItem.OrderID)
	json.NewEncoder(writer).Encode(order_items)

}
