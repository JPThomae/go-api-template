package models

import (
	"time"

	"example.com/rest-api/db"
)

// binding:"required" is used to validate the fields in the struct, it must be passed these values to work
type Event struct {
	Id int
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	// Prepare the SQL statement is optional, but it is a good practice to avoid SQL injection attacks
	statment, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statment.Close()
	result, err := statment.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	// use db.DB.Query to get the rows
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	// use db.DB.QueryRow to get a single row
	statment, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statment.Close()
	row := statment.QueryRow(id)
	
	var event Event
	err = row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	
	return &event, nil
}

func (e Event) UpdateEventById() error{
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?, user_id = ?
	WHERE id = ?`

	// Prepare the SQL statement is optional, but it is a good practice to avoid SQL injection attacks
	statment, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statment.Close()
	_, err = statment.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.Id)

	if err != nil {
		return err
	}
	return nil
}

func (e Event) DeleteEventById(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	// use db.DB.Exec to delete the row
	statment, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statment.Close()
	_, err = statment.Exec(id)

	if err != nil {
		return err
	}
	return nil
}
