package setup

import (
	"github.com/spf13/viper"
	"fmt"
)

type Configuration struct {
	Server		ServerConfig
	Database 	[]DatabaseConfig
	Cache		[]CacheConfig
}

type ServerConfig struct {
	Port		int
	Debug		bool
}

type DatabaseConfig struct {
	Use			string
	Host		string
	Port		int
	User		string
	Password	string
	Database	string
}

type CacheConfig struct {
	Use			string
	Session		int
	Address		string
	Password	string
}

func NewConfig() (Configuration) {
	var config Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return config
}

