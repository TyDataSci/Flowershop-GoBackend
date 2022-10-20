package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemRoutes(t *testing.T) {
	//Testing get of flower id 1
	db.Connect()
	request, _ := http.NewRequest("GET", "/items/1", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}

func TestUserRoutes(t *testing.T) {
	//Testing get of User by username
	db.Connect()
	request, _ := http.NewRequest("GET", "/users/test", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}

func TestSessionRoutes(t *testing.T) {
	//Testing get of Account id 1
	db.Connect()
	request, _ := http.NewRequest("GET", "/user", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}

func TestOrderItemsRoutes(t *testing.T) {
	//Testing get of Order id 1
	db.Connect()
	request, _ := http.NewRequest("GET", "/order_items/order=2", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}
