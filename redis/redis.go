package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

func NewCache(address string, port string, password string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       0,
	})
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(err)
	}

	return &Cache{client: rdb}
}

func (c *Cache) Get(key string) (string, error) {
	value, err := c.client.Get(c.client.Context(), key).Result()
	if err == redis.Nil {
		return "Key does not exist", err
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func (c *Cache) Set(key string, value string, expiration time.Duration) error {
	err := c.client.Set(c.client.Context(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Delete(key string) error {
	err := c.client.Del(c.client.Context(), key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Exists(key string) (bool, error) {
	_, err := c.client.Get(c.client.Context(), key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Cache) GetOrSetDefault(key string, defaultValue string) (string, error) {
	value, err := c.client.Get(c.client.Context(), key).Result()
	if err == redis.Nil {
		err = c.Set(key, defaultValue, 0)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func (c *Cache) Expire(key string, expiration time.Duration) error {
	err := c.client.Expire(c.client.Context(), key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Increment(key string) (int64, error) {
	value, err := c.client.Incr(c.client.Context(), key).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (c *Cache) Decrement(key string) (int64, error) {
	value, err := c.client.Decr(c.client.Context(), key).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (c *Cache) TTL(key string) (time.Duration, error) {
	value, err := c.client.TTL(c.client.Context(), key).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (c *Cache) Keys(pattern string) ([]string, error) {
	keys, err := c.client.Keys(c.client.Context(), pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (c *Cache) FlushAll() error {
	err := c.client.FlushAll(c.client.Context()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Close() error {
	err := c.client.Close()
	if err != nil {
		return err
	}
	return nil
}
