package api

import (
	"Flowershop-GoBackend/pkg/db"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetOrders(writer http.ResponseWriter, request *http.Request) {

}

func CompleteOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	params := mux.Vars(request)
	paramsOrderID, _ := strconv.Atoi(params["id"])
	order, _ := db.GetOrder(paramsOrderID)
	order.Completed = true
	db.UpdateOrder(order)
	c, err := request.Cookie("session-token")
	if err == nil {
		sessionToken := c.Value
		userSession, _ := db.GetUserSession(sessionToken)
		cartOrder, err := db.GetIncompletOrder(userSession.UserID)
		if err != nil {
			cartOrder, err = db.CreateOrder(userSession.UserID)
			fmt.Println("Created order for cart")
			if err != nil {
				fmt.Println("db.CreateOrder", err)
				return
			}
		}

		expiresAt := time.Now().Add(120 * time.Minute)
		userSession.OrderID = cartOrder.ID
		userSession.Expiry = expiresAt
		fmt.Println("Session", userSession.UserID, cartOrder.ID)
		db.UpdateSessionIDs(userSession)

	} else {
		log.Printf("Unable to find session-token for order completed request")
		writer.WriteHeader(401)

	}

}
