package cookie

import (
	"github.com/gorilla/sessions"
)

const ONE_WEEK = 86400 * 7

func NewCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth")
	encryptionKey := []byte("my-encryption")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = ONE_WEEK
	store.Options.HttpOnly = true

	return store
}
