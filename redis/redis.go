package redis

import (
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

const keySession = "session"

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

func Delete(c echo.Context) {
	session := GetSession(c)
	session.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	session.Save(c.Request(), c.Response())
}

func GetCurrentUser(c echo.Context) interface{} {
	store := GetSession(c)
	id := store.Values["userId"]
	return id
}
