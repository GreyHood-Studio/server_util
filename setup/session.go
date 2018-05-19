package setup

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"github.com/GreyHood-Studio/server_util/checker"
)

func createRedisSession(r *gin.Engine, address string, password string) {
	fmt.Println(address, password)
	store, err := redis.NewStore(10, "tcp", address, "", []byte(password))

	checker.CheckError(err, "redis connect checker")

	r.Use(sessions.Sessions("session", store))
}