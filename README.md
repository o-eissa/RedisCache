# Redis Cache

This is a simple Redis cache implementation in Go using the `go-redis/redis` package.

## Installation

To use this package in your Go project, run the following command: go get github.com/your-username/redis-cache


## Usage

To use the Redis cache, first create a new cache instance by calling the `NewCache` function:

```go
cache := redis.NewCache("localhost", "6379", "PASSWORD")

This will create a new Redis client and return a new Cache instance.
```

## API Reference

### `NewCache(address string, port string, password string) *Cache`

Creates a new cache instance with the provided Redis server address, port, and password.

### `Get(key string) (string, error)`

Retrieves the value associated with the given key from the cache.

### `Set(key string, value string, expiration time.Duration) error`

Sets a key-value pair in the cache with the specified expiration time.

### `Delete(key string) error`

Deletes the key from the cache.

### `Exists(key string) (bool, error)`

Checks if the key exists in the cache.

### `GetOrSetDefault(key string, defaultValue string) (string, error)`

Retrieves the value associated with the given key from the cache. If the key does not exist, it sets the key-value pair with the provided default value.

### `Expire(key string, expiration time.Duration) error`

Sets the expiration time for the given key in the cache.

### `Increment(key string) (int64, error)`

Increments the value of the given key in the cache.

### `Decrement(key string) (int64, error)`

Decrements the value of the given key in the cache.

### `TTL(key string) (time.Duration, error)`

Gets the remaining time to live for the given key in the cache.

### `Keys(pattern string) ([]string, error)`

Gets all keys matching the provided pattern from the cache.

### `FlushAll() error`

Flushes all keys from the cache.

### `Close() error`

Closes the cache connection.