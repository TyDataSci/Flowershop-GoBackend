package middleware

import (
	"Flowershop-GoBackend/pkg/models"
	"encoding/gob"

	"github.com/gorilla/sessions"
)

func SessionStore() *sessions.CookieStore {
	var store = sessions.NewCookieStore([]byte("super-secret-sequence"))
	//store.Options.HttpOnly = true
	gob.Register(&models.User{})
	return store
}
