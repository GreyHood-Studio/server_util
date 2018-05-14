package database

import (
	"fmt"
	"database/sql"
)

func init() {
	fmt.Println("initialze db map")
	Conns = make(map[string]*sql.DB)
}
