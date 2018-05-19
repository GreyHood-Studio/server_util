package database

import (
	"github.com/go-redis/redis"
	"time"
)

type UserSession struct {
	conn		*redis.Client
	cacheKey	string
	expired		time.Duration
}
