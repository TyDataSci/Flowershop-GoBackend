package api

import (
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = []*models.User{}

var currentUser = &models.User{ID: "0", Username: "Admin", Password: "Password"}

func GetUsersArray() []*models.User {
	return users
}

func SetUsersArray(_users []*models.User) {
	users = _users
}

func RemoveUser(slice []*models.User, index int) []*models.User {
	modified := make([]*models.User, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func GetCurrentUser() *models.User {
	return currentUser
}

func getUsers(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(users)
}

func getUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for _, user := range users {
		if user.Username == params["username"] {
			json.NewEncoder(writer).Encode(user)
			currentUser = user
			return
		}
	}
}

func createUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var user *models.User
	_ = json.NewDecoder(router.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000))
	users = append(users, user)
	json.NewEncoder(writer).Encode(users)
}

func updateUser(writer http.ResponseWriter, router *http.Request) {
	//Remove the previous user and append the modified user with changes
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, user := range users {
		if user.ID == params["id"] {
			users = RemoveUser(users, index)
			var modifyUser *models.User
			_ = json.NewDecoder(router.Body).Decode(&modifyUser)
			modifyUser.ID = params["id"]
			users = append(users, modifyUser)
			json.NewEncoder(writer).Encode(modifyUser)
		}
	}

}

func deleteUser(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, user := range users {
		if user.ID == params["id"] {
			users = RemoveUser(users, index)
			fmt.Printf("Delete %v\n", user.ID)
			break
		}
	}
}
