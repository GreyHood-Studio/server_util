package database

import (
	"github.com/gin-contrib/sessions"
)

var CacheConn map[string]sessions.RedisStore

func ConnectRedis(use string, session int, address string, password string) {
	CacheConn[use], _ = sessions.NewRedisStore(session, "tcp", address, password)
}