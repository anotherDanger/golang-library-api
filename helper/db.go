package helper

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Db() *sql.DB {

	db, err := sql.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {

		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
