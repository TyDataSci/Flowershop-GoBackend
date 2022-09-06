package api

import (
	"Flowershop-GoBackend/pkg/models"
	"net/http"
)

var userOrderMap = make(map[string][]*models.Order)

func GetUserOrderMap() map[string][]*models.Order {
	return userOrderMap
}

func SetUserOrderMap(_userOrderMap map[string][]*models.Order) {
	userOrderMap = _userOrderMap
}

func RemoveOrder(slice []*models.Order, index int) []*models.Order {
	modified := make([]*models.Order, 0)
	modified = append(modified, slice[:index]...)
	return append(modified, slice[index+1:]...)
}

func GetOrders(writer http.ResponseWriter, router *http.Request) {

}

func GetOrder(writer http.ResponseWriter, router *http.Request) {

}
