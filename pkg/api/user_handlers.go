package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/middleware"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	fmt.Println("params ", params["username"])
	user, err := db.GetUser(params["username"])
	fmt.Println(err)
	if err == nil {
		json.NewEncoder(writer).Encode(user)
		cookie := &http.Cookie{Name: "my-cookie", Value: "sessionToken"}
		http.SetCookie(writer, cookie)
	}
}

func CreateUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "https://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")
	var newUser models.User
	json.NewDecoder(router.Body).Decode(&newUser)
	bytes, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(bytes)
	newUser, err := db.CreateUser(newUser.Username, newUser.Name, newUser.Password)
	if err == nil {
		json.NewEncoder(writer).Encode(newUser)
	}
}

func ValidateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "https://foreveryoursflowershop.com")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cookie")

	var validateUser models.User
	json.NewDecoder(request.Body).Decode(&validateUser)
	user, err := db.GetUser(validateUser.Username)
	if err == nil {
		err := middleware.HashwordCompare(user.Password, validateUser)
		if err == nil {
			json.NewEncoder(writer).Encode(user)
			cartOrder, err := db.GetIncompletOrder(user.ID)
			if err != nil {
				cartOrder, err = db.CreateOrder(user.ID)
				fmt.Println("Created order for cart")
				if err != nil {
					fmt.Println("db.CreateOrder", err)
					return
				}
			}
			c, err := request.Cookie("session-token")
			if err == nil {
				sessionToken := c.Value
				expiresAt := time.Now().Add(120 * time.Minute)
				userSession := models.Session{Token: sessionToken, UserID: user.ID, OrderID: cartOrder.ID, Expiry: expiresAt}
				fmt.Println("Session", user.ID, cartOrder.ID)
				db.UpdateSessionIDs(userSession)

			}

		} else {
			log.Printf("Incorrect Username or Password for %v", validateUser.Username)
			writer.WriteHeader(401)

		}
	}
}
