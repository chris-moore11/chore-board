package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "1234"
	dbname   = "db"
)

// Global DB
var GlobalInstance *sql.DB

func ConnectDB() {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		// Could not connect to db, so attempt to create new db
		db = InitDB()
	}

	fmt.Printf("\nSuccessfully connected to database!\n")
	GlobalInstance = db
}

// Creates database and populates with starting values. Requires existence of "postgres" user with password "1234"
func InitDB() *sql.DB {
	fmt.Println("Attempting to create database " + dbname)
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable", user, password, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	ExecOrPanic(db, "create database "+dbname)
	fmt.Println("Database " + dbname + " created")

	fmt.Println("Connecting to database " + dbname)
	connStr = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err = sql.Open("postgres", connStr)

	fmt.Println("Creating users table")
	ExecOrPanic(db, `
		CREATE TABLE users (  
			id INTEGER PRIMARY KEY NOT NULL,  
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			demerits INTEGER NOT NULL,
			choreId INTEGER NOT NULL,
			admin BOOLEAN NOT NULL
		);
	`)
	ExecOrPanic(db, `
		INSERT INTO users (id, name, email, demerits, choreId, admin)
		VALUES 
		(1, 'Anthony', 'anthony@gmail.com', 0, 1, false),
		(2, 'Michael', 'michael@gmail.com', 0, 1, false),
		(3, 'Frits', 'frits@gmail.com', 0, 2, false),
		(4, 'Conor', 'conor@gmail.com', 0, 3, false),
		(5, 'Noah', 'noah@gmail.com', 0, 4, true),
		(6, 'Chris', 'chris@gmail.com', 0, 5, true),
		(7, 'Bart', 'bart@gmail.com', 0, 5, false);
	`)

	fmt.Println("Creating chores table")
	ExecOrPanic(db, `
		CREATE TABLE chores (  
			id INTEGER PRIMARY KEY NOT NULL,  
			text TEXT NOT NULL,
			done BOOLEAN NOT NULL,
			description TEXT NOT NULL,
			image TEXT NOT NULL
		);
	`)
	ExecOrPanic(db, `
		INSERT INTO chores (id, text, done, description, image)
		VALUES
		(1, 'Clean Kitchen', false, '', ''),
		(2, 'Clean Floors', false, '', ''),
		(3, 'Trash', false, '', ''),
		(4, 'CHLORD', false, '', ''),
		(5, 'Off', false, '', '');
	`)

	fmt.Printf("\nSuccessfully initialized database %s!\n", dbname)

	return db
}

func ExecOrPanic(db *sql.DB, command string) {
	_, err := db.Exec(command)
	if err != nil {
		panic(err)
	}
}
