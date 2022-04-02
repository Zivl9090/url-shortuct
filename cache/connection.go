package cache

import (
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func Connect() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
