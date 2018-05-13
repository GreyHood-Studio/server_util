package database

import (
	"github.com/gin-contrib/sessions"
	"time"
)

var CacheConn map[string]sessions.RedisStore
//var Sessions
//var caches

func ConnectRedis(use string, session int, address string, password string) {
	CacheConn[use], _ = sessions.NewRedisStore(session, "tcp", address, password)
}

func CloseRedis()  {

}

func CreateRedisSession(session int, host string, password string, defaultExpiration time.Duration) {
	
}

func CreateRedisCache()  {
	
}