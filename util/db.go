package util

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" //import driver
)

//NewDBClient to create new connection
func NewDBClient() (*sql.DB, error) {
	var dbURI string
	dbURI = os.Getenv("CONNECTION_STRING")

	// dbPool is the pool of database connections.
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
