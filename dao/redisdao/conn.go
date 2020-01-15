package redisdao

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var client *redis.Client

// Init build redis connection
func Init() {
	host := viper.GetString("redis.host")
	options := redis.Options{}
	options.Addr = host
	client = redis.NewClient(&options)
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
