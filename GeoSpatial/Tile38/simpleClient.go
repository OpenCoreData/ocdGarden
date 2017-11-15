package main

import (
	"log"

	redis "gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:9851",
	})

	// cmd := redis.NewStringCmd("SET", "fleet", "truck1", "POINT", 23.32, 115.423)
	// client.Process(cmd)
	// v, _ := cmd.Result()
	// log.Println(v)

	// cmd1 := redis.NewStringCmd("GET", "test", "hono")
	cmd1 := redis.NewStringCmd("GET", "fleet", "truck3")
	client.Process(cmd1)
	v1, _ := cmd1.Result()
	log.Println(v1)

	client.Close()
}
