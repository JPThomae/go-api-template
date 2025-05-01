package db

import (
	"fmt"
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Println(DB.Stats())

	createEventsTables()
	createUserTable()
	createRegistrationTable()
}


func createEventsTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(fmt.Sprintf("could not create events table: %v", err))
	}
}

func createUserTable() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic(fmt.Sprintf("could not create users table: %v", err))
	}
}

func createRegistrationTable() {
	createRegistration := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id, INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err := DB.Exec(createRegistration)

	if err != nil {
		panic(fmt.Sprintf("could not create registration table: %v", err))
	}
}
