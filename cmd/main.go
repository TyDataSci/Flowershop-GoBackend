package main

import (
	"Flowershop-GoBackend/pkg/api"
	"Flowershop-GoBackend/pkg/dev"
	"Flowershop-GoBackend/pkg/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//initialize router to handle api calls
	dev.InitializeMockData()
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", api.Router()))

	// Stores session using secure cookies
	//*CookieStore  --> Struct with codecs to store cookies and options

	var c *gin.Context
	session, _ := middleware.SessionStore().Get(c.Request, "session")
	session.Values["UserID"] = "admin"
	session.Save(c.Request, c.Writer)

}
