package models

import (
	"errors"

	"kodski.com/events-api/db"
	"kodski.com/events-api/utils"
)

type User struct {
	ID int64
	Email string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) Authenticate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var hashedPass string
	err := row.Scan(&u.ID, &hashedPass)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, hashedPass)

	if !passwordIsValid {
		return errors.New("invalid details")
	}
	return nil
}