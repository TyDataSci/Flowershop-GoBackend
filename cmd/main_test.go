package main

import (
	"Flowershop-GoBackend/pkg/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowerRoutes(t *testing.T) {
	//Testing get of flower id 1
	request, _ := http.NewRequest("GET", "/flowers/1", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}

func TestUserRoutes(t *testing.T) {
	//Testing get of User by username
	request, _ := http.NewRequest("GET", "/users/tssand", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}

func TestAccountsRoutes(t *testing.T) {
	//Testing get of Account id 1
	request, _ := http.NewRequest("GET", "/accounts/1", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}

func TestOrdersRoutes(t *testing.T) {
	//Testing get of Order id 1
	request, _ := http.NewRequest("GET", "/orders/1", nil)
	response := httptest.NewRecorder()
	//The response recorder used to record HTTP responses
	api.Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK Response expected")
}
