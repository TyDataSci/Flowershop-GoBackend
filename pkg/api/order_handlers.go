package api

import (
	"Flowershop-GoBackend/pkg/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOrders(writer http.ResponseWriter, request *http.Request) {

}

func CompleteOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "https://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	params := mux.Vars(request)
	paramsUserID, _ := strconv.Atoi(params["id"])
	order, _ := db.GetIncompletOrder(paramsUserID)
	order.Completed = true
	err := db.UpdateOrder(order)
	if err == nil {
		session, _ := db.GetUserLastSession(paramsUserID)
		order, _ := db.CreateOrder(paramsUserID)
		session.OrderID = order.ID
		db.UpdateSessionIDs(session)
	} else {
		log.Printf("Unable to find session-token for order completed request")
		writer.WriteHeader(401)

	}

}
