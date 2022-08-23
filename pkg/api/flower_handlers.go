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

var flowers = []*models.Item{}

func RemoveItem(slice []*models.Item, index int) []*models.Item {
	modified := make([]*models.Item, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func getFlowers(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(writer).Encode(flowers)
}

func getFlower(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for _, item := range flowers {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func createFlower(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var flower *models.Item
	_ = json.NewDecoder(router.Body).Decode(&flower)
	flower.ID = strconv.Itoa(rand.Intn(100000))
	flowers = append(flowers, flower)
	json.NewEncoder(writer).Encode(flowers)
}

func updateFlower(writer http.ResponseWriter, router *http.Request) {
	//set json content type
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	//json params
	params := mux.Vars(router)
	//loop thru movies,range
	for index, item := range flowers {
		if item.ID == params["id"] {
			flowers = RemoveItem(flowers, index)
			//flowers = append(flowers[:index], flowers[index+1:]...)
			var newFlower *models.Item
			_ = json.NewDecoder(router.Body).Decode(&newFlower)
			newFlower.ID = params["id"]
			flowers = append(flowers, newFlower)
			json.NewEncoder(writer).Encode(newFlower)
		}
	}
}

func deleteFlower(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(router)
	for index, item := range flowers {
		if item.ID == params["id"] {
			flowers = RemoveItem(flowers, index)
			fmt.Printf("Deleted %v\n", item.ID)
			break
		}
	}
	json.NewEncoder(writer).Encode(flowers)
}
