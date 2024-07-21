package models

import "kodski.com/events-api/db"

type Registration struct {
	ID int64 `json:"id"`
	EventID int64 `binding:"required" json:"eventId"`
	UserID int64 `binding:"required" json:"userId"`
}

func (r *Registration) Save() error {
	query := `
	INSERT INTO registrations (eventId, userId)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}
	registrationId, err := result.LastInsertId()
	r.ID = registrationId
	return err
}

func DeleteRegistration(eventId, userId int64) error {
	query := `
	DELETE FROM registrations
	WHERE eventId = ? AND userId = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(eventId, userId)
	return err
}