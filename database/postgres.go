package database


import (
	_ "github.com/lib/pq"
	"database/sql"
)

func ConnectPG(use string, psql string) *sql.DB{
	conn, err := sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	return conn
}