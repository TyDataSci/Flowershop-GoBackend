package api

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GetUserSession(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://foreveryoursflowershop.com")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c, err := request.Cookie("session-token")
	if err == nil {
		sessionToken := c.Value
		userSession, _ := db.GetUserSession(sessionToken)
		json.NewEncoder(writer).Encode(userSession)
		fmt.Println("Returned Session")
		println("Got same token finished")

	} else {
		sessionToken := uuid.NewString()
		expiresAt := time.Now().Add(120 * time.Minute)
		userSession := models.Session{Token: sessionToken, UserID: 0, OrderID: 0, Expiry: expiresAt}
		userSession, err := db.CreateSession(userSession)

		if err != nil {
			fmt.Println("db.CreateSession", err)
			return
		}
		cookieSession := &http.Cookie{Name: "session-token", Value: userSession.Token, Expires: userSession.Expiry, Path: "/", Domain: ".foreveryoursflowershop.com", SameSite: http.SameSiteNoneMode,
			Secure: true}
		http.SetCookie(writer, cookieSession)
		println("Created Token finished")
		json.NewEncoder(writer).Encode(userSession)

	}

}
