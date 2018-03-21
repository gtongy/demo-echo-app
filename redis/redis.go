package redis

import (
	redistore "gopkg.in/boj/redistore.v1"
)

func Init() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	return store
}
