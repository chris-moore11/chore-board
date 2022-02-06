package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "1234"
    dbname = "db"
)

// Global DB
var GlobalInstance *sql.DB

func ConnectDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if (err != nil) {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	GlobalInstance = db
}
