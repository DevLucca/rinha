package cache

import (
	"errors"
)

var ErrCacheMiss = errors.New("cache miss: key not found")

type Cache interface {
	GetItem(key string) ([]byte, error)
	SetItem(key string, data []byte) error
	Delete(keys ...string) error
	SetString(key string, data string) error
	GetString(key string) (string, error)
	GetInt(key string) (int64, error)
	Increase(key string) error
}
