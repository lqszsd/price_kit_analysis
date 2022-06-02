package utils

import (
	"github.com/go-redis/redis/v8"
)

var re *redis.Client

func GetRedis() *redis.Client {
	if re == nil {
		re = redis.NewClient(&redis.Options{
			Addr:     "114.116.77.103:6379",
			Password: "lq123456789", // no password set
			DB:       0,  // use default DB
		})
	}
	return re
}

