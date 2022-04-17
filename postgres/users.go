package database

import (
	"fmt"
	"time"

	"github.com/DavidHODs/TechHUB-goGraph/customerror"
)

// models the user details
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


// It saves the registered user details into the database
func SaveUser() (int64, error) {
	var user User
	stmt, err := Db.Prepare("INSERT INTO tech.users(name, email, password) VALUES(?, ?, ?")
	if err != nil {
		customerror.HandleError(err, true)
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		customerror.HandleError(err, true)
	}

	id, err := res.LastInsertId()
	if err != nil {
		customerror.HandleError(err, true)
	}

	rows, _ := res.RowsAffected()

	fmt.Printf("%d rows affected", rows)

	return id, err
}