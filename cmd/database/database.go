package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(credentials string) error {
	var err error
	DB, err = sql.Open("mysql", credentials)

	if err != nil {
		return fmt.Errorf("could not connect to DB, %v", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return nil
}
