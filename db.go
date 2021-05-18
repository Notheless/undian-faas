package p

import (
	"database/sql"
	"time"
)

//NewDBClient to create new connection
func NewDBClient() *sql.DB {

	db, err := sql.Open("mysql", "root:3Au6zdCNx6x1wLhg@34.101.149.129/default")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
