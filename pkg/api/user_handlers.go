package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RemoveUser(username string) {
}

func GetUsers(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	//json.NewEncoder(writer).Encode(users)
}

func GetUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	user, err := db.GetUser(params["username"])
	if err == nil {
		json.NewEncoder(writer).Encode(user)
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

func UpdateUser(writer http.ResponseWriter, router *http.Request) {
}

func DeleteUser(writer http.ResponseWriter, router *http.Request) {
}
