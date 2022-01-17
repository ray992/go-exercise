package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := rdb.Ping(ctx).Result()
	fmt.Println(err)

	setErr := rdb.Set(ctx, "a", 0, 0).Err()
	if setErr != nil {
		fmt.Println(setErr)
		return
	}
	res , err := rdb.Get(ctx, "a").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
