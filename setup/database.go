package setup

import (
	"fmt"
	"github.com/GreyHood-Studio/server_util/database"
)

func ConnectDatabase(dbs []DatabaseConfig){
	// rdb
	setupPostgres(dbs)
}

func setupPostgres(configs []DatabaseConfig) {
	var psql string
	for _, config := range configs {
		if config.Use != "game" && config.Use != "user" {
			//checker
			return
		}
		psql = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Database)
		database.ConnectPG(config.Use, psql)
		fmt.Println(config.Use, "db Successfully connected!")
	}
}
