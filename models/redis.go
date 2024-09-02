package models

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx context.Context

func RedisInit() {
	Ctx = context.Background()
	Client = NewClient(Ctx)
}

func NewClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if pong, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}

	return client
}
