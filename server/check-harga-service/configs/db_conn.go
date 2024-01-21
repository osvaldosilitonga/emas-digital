package configs

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "postgres://adminpostgres:87654321@localhost:5435/emasdigital?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Println("Connection Fail")
	} else {
		log.Println("DB Connected Successfully")
	}

	return db
}
