package models

import (
	"github.com/go-redis/redis"
)

//Declaring a redis client
var client *redis.Client

func Init() {
	//Initiate a new Redis client
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

}
