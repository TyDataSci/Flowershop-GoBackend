package api

import (
	"Flowershop-GoBackend/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func GetOrderItems(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	params := mux.Vars(router)
	fmt.Println(params)
	orderid, _ := strconv.Atoi(params["orderid"])
	fmt.Printf("orderid %v \n", orderid)
	order_items, _ := db.GetOrderItems(orderid)
	json.NewEncoder(writer).Encode(order_items)
}

func CreateOrderItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	params := mux.Vars(router)
	fmt.Println(params)
	fmt.Println("Create function")
	paramsOrderID, _ := strconv.Atoi(params["orderid"])
	paramsItemID, _ := strconv.Atoi(params["itemid"])
	fmt.Printf("paramOrderID %v\n", paramsOrderID)
	fmt.Printf("paramItemID %v\n", paramsOrderID)
	db.CreateOrderItem(paramsOrderID, paramsItemID)
	order_items, _ := db.GetOrderItems(paramsOrderID)
	json.NewEncoder(writer).Encode(order_items)
}

func RemoveOrderItem(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	params := mux.Vars(router)
	fmt.Println(params)
	fmt.Println("Remove function")
	paramsOrderID, _ := strconv.Atoi(params["orderid"])
	paramsItemID, _ := strconv.Atoi(params["itemid"])
	paramRemoved := true
	db.UpdateOrderItem(paramsOrderID, paramsItemID, paramRemoved)
	order_items, _ := db.GetOrderItems(paramsOrderID)
	json.NewEncoder(writer).Encode(order_items)

}
