package main

import (
	"github.com/o-eissa/redis-cache/redis"
	"github.com/o-eissa/redis-cache/utils"
)

func main() {
	// Example usage of the rediscache package
	address, port, password := utils.LoadEnvVars()

	cache := redis.NewCache(address, port, password)

	cache.Increment("counter")
	cache.Increment("counter")
	cache.Set("username", "o-eissa", 0)

	value, err := cache.Get("counter")
	if err != nil {
		panic(err)
	}
	println("Counter: " + value)

	value, err = cache.Get("username")
	if err != nil {
		panic(err)
	}
	println("Username: " + value)

}
