package redis

import (
	"fmt"
	"strconv"

	"github.com/DevLucca/rinha/infra/persistence/cache"

	"gopkg.in/redis.v5"
)

type Client struct {
	redis  *redis.Client
	prefix string
}

type ConfigOptions struct {
	Server     string
	DB         int
	Password   string
	Port       int
	Prefix     string
	Expiration string
}

func NewClient(configs ConfigOptions) cache.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", configs.Server, configs.Port),
		Password: configs.Password,
		DB:       configs.DB,
	})

	return &Client{
		redis:  client,
		prefix: configs.Prefix,
	}
}

func (c *Client) buildKey(key string) string {
	if c.prefix != "" {
		return c.prefix + ":" + key
	}

	return key
}

func (c *Client) GetItem(key string) ([]byte, error) {
	val, err := c.redis.Get(c.buildKey(key)).Bytes()
	if err == redis.Nil {
		return val, cache.ErrCacheMiss
	} else if err != nil {
		return val, fmt.Errorf("error cache get item: %w", err)
	}

	return val, nil
}

func (c *Client) SetItem(key string, data []byte) error {
	err := c.redis.Set(c.buildKey(key), data, 0).Err()
	if err != nil {
		return fmt.Errorf("error cache set item: %w", err)
	}

	return nil
}

func (c *Client) Delete(keys ...string) error {
	redisKeys := make([]string, len(keys))
	for i, key := range keys {
		redisKeys[i] = c.buildKey(key)
	}
	err := c.redis.Del(redisKeys...).Err()
	if err != nil {
		return fmt.Errorf("error cache delete: %w", err)
	}

	return nil
}

func (c *Client) SetString(key string, data string) error {
	return c.SetItem(key, []byte(data))
}

func (c *Client) GetString(key string) (string, error) {
	val, err := c.GetItem(key)
	if err == cache.ErrCacheMiss {
		return "", cache.ErrCacheMiss
	} else if err != nil {
		return "", fmt.Errorf("error cache get string: %w", err)
	}

	return string(val), nil
}
func (c *Client) GetInt(key string) (int64, error) {
	val, err := c.GetItem(key)
	if err == cache.ErrCacheMiss {
		return 0, cache.ErrCacheMiss
	} else if err != nil {
		return 0, fmt.Errorf("error cache get int: %w", err)
	}

	return strconv.ParseInt(string(val), 10, 64)
}

func (c *Client) SetInt(key string, data int64) error {
	err := c.redis.Set(c.buildKey(key), data, 0).Err()
	if err != nil {
		return fmt.Errorf("error cache set int: %w", err)
	}

	return nil
}

func (c *Client) Increase(key string) error {
	err := c.redis.Incr(c.buildKey(key)).Err()
	if err != nil && err == redis.Nil {
		return cache.ErrCacheMiss
	} else if err != nil {
		return fmt.Errorf("error cache to increase %w", err)
	}

	return nil
}

func (c *Client) Exists(key string) bool {
	_, err := c.GetItem(key)
	if err == cache.ErrCacheMiss {
		return false
	} else if err != nil {
		return false
	}

	return true
}
