package redis

import (
	"os"

	"github.com/boj/redistore"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/soveran/redisurl"
)

const (
	keySession     = "session"
	defaultAddress = "tcp://redis:6379"
)

var (
	redisPool *redis.Pool
)

func init() {
	url := address()
	redisPool = redis.NewPool(func() (redis.Conn, error) {
		return redisurl.ConnectToURL(url)
	}, 10)
}

func GetStore() *redistore.RediStore {
	store, err := redistore.NewRediStoreWithPool(redisPool, []byte("secret-key"))
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

func address() string {
	address := os.Getenv("REDISTOGO_URL")
	if address == "" {
		return defaultAddress
	}
	return address
}
