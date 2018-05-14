package setup

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/GreyHood-Studio/server_util/database"
	"github.com/GreyHood-Studio/server_util/error"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
)

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

func setupPostgres(configs []DatabaseConfig) {
	var psql string
	for _, config := range configs {
		if config.Use != "game" && config.Use != "user" {
			//error
			return
		}
		psql = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Database)
		database.ConnectPG(config.Use, psql)
		fmt.Println(config.Use, "db Successfully connected!")
	}
}

func ConnectRouter(caches []CacheConfig) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// cache layer
	for _, config := range caches {
		if config.Use == "cache" {
			fmt.Println(config.Use, "redis Successfully connected!")
		} else if config.Use == "session" {
			createRedisSession (r, config.Address, config.Password)
			fmt.Println(config.Use, "redis Successfully connected!")
		}
	}

	// setup route 설정 전에 넘겨야함
	return r
}

func createRedisSession(r *gin.Engine, address string, password string) {
	fmt.Println(address, password)
	store, err := redis.NewStore(10, "tcp", address, "", []byte(password))

	error.CheckError(err, "redis connect error")

	r.Use(sessions.Sessions("session", store))
}

func ConnectDatabase(dbs []DatabaseConfig){
	// rdb
	setupPostgres(dbs)
}
