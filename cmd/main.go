package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db.Connect()
	//dev.InitializeMockData()
	//db.CreateUser("Admin", "Admin", "Password")
	//user, err := db.CreateUser("Admin", "Admin", "Password")
	//if err == nil {
	//		order, err := db.CreateOrder(user.ID)
	//		if err == nil {
	//			db.CreateOrderItem(order.ID, 1)
	//			db.CreateOrderItem(order.ID, 2)
	//			db.CreateOrderItem(order.ID, 3)
	//			items, _ := db.GetItems()
	//			for _, item := range items {
	//				println(item.Description)
	//			}
	//	db.GetOrder(order.UserID)
	//	order.Delivery = true
	//	order.Completed = true
	//	db.UpdateOrder(order)

	//		}

	//	}
	//initialize router to handle api calls
	//dev.InitializeMockData()
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", api.Router()))

	// Stores session using secure cookies
	//*CookieStore  --> Struct with codecs to store cookies and options

	var c *gin.Context
	session, _ := middleware.SessionStore().Get(c.Request, "session")
	session.Values["UserID"] = "admin"
	session.Save(c.Request, c.Writer)

}
