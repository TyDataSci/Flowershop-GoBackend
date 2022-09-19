package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var newUser models.User
	json.NewDecoder(router.Body).Decode(&newUser)
	newUser, err := db.CreateUser(newUser.Username, newUser.Name, newUser.Password)
	if err == nil {
		json.NewEncoder(writer).Encode(newUser)
	}
}

func ValidateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	var validateUser models.User
	json.NewDecoder(request.Body).Decode(&validateUser)
	user, err := db.GetUser(validateUser.Username)
	if err == nil {
		if validateUser.Password == user.Password {
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
