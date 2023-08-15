package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// Profile configuration profile
type Profile string

func (p Profile) String() string {
	return string(p)
}

// Profile default values
const (
	ProfileRun Profile = "config"
)

// EnvKeyReplacer replace for environment variable parse
var EnvKeyReplacer = strings.NewReplacer(".", "_", "-", "_")

func setup() {
	viper.SetEnvKeyReplacer(EnvKeyReplacer)
	viper.AutomaticEnv()

	env := os.Getenv("ENV")

	if env == "" {
		env = "dev"
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(path.Dir(os.Args[0]))
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
}

// Read returns the configuration values,
//
//	based on the configuration files and environment variables.
func Read() (cfg *Config, err error) {
	setup()

	fillDefault()
	viper.ReadInConfig()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
