package redis

import (
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

const keySession = "session"

var (
	store *sessions.Session
)

func Init() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	return store
}

func GetSession(c echo.Context) *sessions.Session {
	store, err := session.Get(keySession, c)
	if err != nil {
		panic(err)
	}
	return store
}

func GetCurrentUser(c echo.Context) interface{} {
	session := GetSession(c)
	id := session.Values["userId"]
	return id
}
