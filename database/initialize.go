package database

import (
	"fmt"
	"database/sql"
	"github.com/gin-contrib/sessions"
)

func init() {
	fmt.Println("initialze db map")
	DBConn = make(map[string]*sql.DB)
	CacheConn = make(map[string]sessions.RedisStore)
}
