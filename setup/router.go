package setup

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"github.com/gin-contrib/zap"
)

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

	logger := setLogger()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// setup route 설정 전에 넘겨야함
	return r
}
