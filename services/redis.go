package services

import (
	"github.com/go-redis/redis"
)

func SetString(key string, value string) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetString(key string) string {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}
