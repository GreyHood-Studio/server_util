package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Conns map[string]*sql.DB

func ConnectPG(use string, psql string) {
	var err error
	Conns[use], err = sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}

	err = Conns[use].Ping()
	if err != nil {
		panic(err)
	}

}

func ClosePG(use string) {
	Conns[use].Close()
	delete(Conns, use)
}