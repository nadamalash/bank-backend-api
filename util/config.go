package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // path to look for the config file in
	viper.SetConfigName("app") // name of the config file without extension
	viper.SetConfigType("env") // type of the config file (json, yaml, toml, env, etc)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal means to convert or decode JSON into a Go data structure
	err = viper.Unmarshal(&config) // unmarshal the config file into the config struct, which is passed by reference
	return
}
