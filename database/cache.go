package database

import (
	"github.com/go-redis/redis"
	"github.com/GreyHood-Studio/server_util/checker"
)

// 서버가 사용하는 캐시
type ServerCache struct {
	conn		*redis.Client
	cacheKey	string
}

func (cache *ServerCache) SetKey(key string, value interface{}){
	// cache key +
	err := cache.conn.Set(key, value, 0).Err()
	checker.NoDeadError(err, "set key error")
}

func (cache *ServerCache) GetKey(key string) interface{}{
	val, err := cache.conn.Get("key").Result()
	checker.NoDeadError(err, "get key error")
	if err == redis.Nil {
		checker.NoDeadError(err, "key not exist")
		return nil
	}
	return val
}

func NewServerCache(address string, password string) *ServerCache{
	serverCache := &ServerCache{
		conn: ConnectRedis(address, password),
	}
	return serverCache
}