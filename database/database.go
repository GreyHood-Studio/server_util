package database

import "database/sql"

// userdb는 오직 webserver만 사용
type UserDB struct {

}

type GameDB struct {
	conn		*sql.DB
}

//func (db *GameDB) {
//
//}