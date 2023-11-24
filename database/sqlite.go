package database

import (
	"context"
	"database/sql"
	"log"
	"time"
)

// ISO8601: YYYY-MM-DD HH:MM:SS.SSS
// const databaseDateTimeFormat = "2006-01-02 15:04:05.000"

func NewSQLite() *sql.DB {
	db, err := sql.Open("sqlite3", "./school.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	return db
}
