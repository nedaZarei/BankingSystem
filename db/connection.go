package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "postgres"
	port     = 5432
	user     = "neda.z"
	password = "nz2003nz"
	dbname   = "banking_system"
)

func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
