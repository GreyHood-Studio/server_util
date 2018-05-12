package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DBConn map[string]*sql.DB

func ConnectPG(use string, psql string) {
	var err error
	DBConn[use], err = sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}

	err = DBConn[use].Ping()
	if err != nil {
		panic(err)
	}

}

func ClosePG(use string) {
	DBConn[use].Close()
	delete(DBConn, use)
}