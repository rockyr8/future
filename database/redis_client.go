package database

import (
	// "fmt"

	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var once sync.Once
var RedisClient *redis.Client

func init() {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:         "45.63.62.104:6379",
			Password:     "Rocky89226will",
			MaxRetries:   3,
			DialTimeout:  5 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolSize:     2000,
			PoolTimeout:  0,
			IdleTimeout:  0,
			DB:           3,
		})

		pong, err := client.Ping().Result()
		log.Println(pong, err)
		RedisClient = client
	})
}

func RedisGet(key string) string {
	var rat string
	val, err := RedisClient.Get(key).Result()
	if err == redis.Nil {
		rat = ""
	} else if err != nil {
		rat = ""
		log.Println(err)
	} else {
		rat = val
	}
	return rat
}

func RedisSet(key string, val string, sec time.Duration) error {
	err := RedisClient.Set(key, val, sec*time.Second).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil	
}

func RedisDel(keys ...string) error {
	err := RedisClient.Del(keys...).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

