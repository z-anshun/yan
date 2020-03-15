package second_work

import (
	"github.com/go-redis/redis"
)

var R *redis.Client

var host = "127.0.0.1:6379"

func init() {
	R = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	R.Ping()
}
