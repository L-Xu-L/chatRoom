package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var (
	Connection *redis.Client
)

const (
	DB = 0
	PASSWORD  = ""
	HOST   = "127.0.0.1"
	PORT     =  "6379"
)

func NewConnetion() {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s",HOST,PORT),
		Password: PASSWORD, // no password set
		DB:       DB,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(pong, err)
	}

	Connection = client
}