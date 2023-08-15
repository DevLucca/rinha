package config

import (
	"github.com/spf13/viper"
)

type (
	// Config represents the configuration values for the whole service.
	Config struct {
		Db    DbConfig    `mapstructure:"db"`
		Cache CacheConfig `mapstructure:"cache"`
	}

	DbConfig struct {
		InMemory bool `mapstructure:"in-memory"`

		Host string `mapstructure:"host"`
		Port uint16 `mapstructure:"port"`
		Name string `mapstructure:"name"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
	}

	CacheConfig struct {
		Server     string `mapstructure:"server"`
		DB         int    `mapstructure:"db"`
		Password   string `mapstructure:"pass"`
		Port       int    `mapstructure:"port"`
		Prefix     string `mapstructure:"prefix"`
		Expiration string `mapstructure:"expiration"`
	}
)

func fillDefault() {
	viper.SetDefault("cache.port", 6379)
	viper.SetDefault("cache.server", "localhost")
	viper.SetDefault("cache.db", 0)
	viper.SetDefault("cache.prefix", "")

	viper.SetDefault("db.in-memory", false)
	viper.SetDefault("db.host", "")
	viper.SetDefault("db.port", "")
	viper.SetDefault("db.name", "")
	viper.SetDefault("db.user", "")
	viper.SetDefault("db.pass", "")
}
