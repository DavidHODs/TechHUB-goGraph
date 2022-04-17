package database

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserMethods interface {
	SavePost() (int64, error) 
}

func SavePost() (int64, error) {
	var user User
	stmt, err := Db.Prepare("INSERT INTO tech.users(name, email, password) VALUES(?, ?, ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rows, _ := res.RowsAffected()

	fmt.Printf("%d rows affected", rows)

	return id, err
}