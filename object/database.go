package object

import "github.com/jmoiron/sqlx"

var Db *sqlx.DB

func RegisterDb(db *sqlx.DB){
	Db = db
}
