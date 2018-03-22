package redis

import (
	"net/http"

	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
)

const keySession = "sessions"

var (
	store *redistore.RediStore
)

func Init() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	return store
}

func GetSession(h *http.Request) *sessions.Session {
	session, err := store.Get(h, keySession)
	if err != nil {
		panic(err)
	}
	return session
}
