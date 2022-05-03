package database

import (
	"database/sql"
	"errors"

	"github.com/DavidHODs/TechHUB-goGraph/utils"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

// It saves the registered user details into the database
func SaveUser(name, email, password, passwordConfirmation string) (string, []byte, error) {
	passwordError := utils.PasswordCheck(password, passwordConfirmation)
	if passwordError != nil {
		utils.HandleError(passwordError, false)
		return "", nil, passwordError
	}

	stmt, err := Db.Prepare(`INSERT INTO tech.users(name, email, password) 
							VALUES($1, $2, $3)
							RETURNING id`)
	if err != nil {
		utils.HandleError(err, false)
	}

	defer stmt.Close()

	var id string = ""

	hashedP, _ := utils.HashPassword(password)

	err = stmt.QueryRow(name, email, hashedP).Scan(&id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == pgerrcode.UniqueViolation {
				utils.HandleError(utils.DupError(email), false)
				pqError := &pq.Error{Message: error.Error(utils.DupError(email))}
				
				return id, nil, pqError
			}
		} 
			utils.HandleError(err, false)
	}

	return id, hashedP, err
}

// Returns user details. Limited to id, name and email for now as well as error if any
func ReturnUserDetails(userId string) (string, string, string, error) {
	stmt, err := Db.Prepare(`SELECT id, name, email FROM tech.users where id = $1`)
	if err != nil {
		utils.HandleError(err, false)
		return "", "", "", errors.New("something went wrong, try again later")
	}

	defer stmt.Close()

	var (
		id string = ""
		name string = ""
		email string = ""
	)

	err = stmt.QueryRow(userId).Scan(&id, &name, &email)
	if err != nil {
		utils.HandleError(err, false)
		return "", "", "", errors.New("something went wrong, try again later")
	}

	return id, name, email, nil
}

// returns the id of user by checking via supplied email in the database
func GetUserIdByEmail(email string) (string, error) {
	stmt, err := Db.Prepare(`SELECT id from tech.users where email = $1`)
	if err != nil {
		utils.HandleError(err, false)
		return "", err
	}

	defer stmt.Close()

	var id string = ""

	err = stmt.QueryRow(email).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.HandleError(errors.New("user does not exist"), false)
			return "", errors.New("user does not exist")
		}

		return "", err
	}

	return id, nil
}