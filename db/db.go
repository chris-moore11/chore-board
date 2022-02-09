package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
    host = "database"
    port = "5432"
    user = "pi"
    password = "1234"
    dbname = "db"
)

// Global DB
var GlobalInstance *sql.DB

func ConnectDB() {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", 
		user, password, host, port, dbname)
	fmt.Sprintf(connStr)
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
