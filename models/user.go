package models

import (
	"github.com/PrinceM13/go-event-booking/db"
	"github.com/PrinceM13/go-event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id
	u.Password = ""

	return err
}
