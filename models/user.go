package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utilities"
)

type User struct {
	Id int
	Username string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utilities.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Username, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = int(userId)
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	var hashedPassword string

	err = statement.QueryRow(u.Email).Scan(&u.Id, &hashedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utilities.CheckPasswordHash(u.Password, hashedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
