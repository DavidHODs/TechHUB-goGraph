package database

import (
	"github.com/DavidHODs/TechHUB-goGraph/utils"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

// It saves the registered user details into the database
func SaveUser(name, email, password, passwordConfirmation string) (int64, error) {
	passwordError := utils.PasswordCheck(password, passwordConfirmation)
	if passwordError != nil {
		utils.HandleError(passwordError, false)
		return 0, passwordError
	}

	stmt, err := Db.Prepare(`INSERT INTO tech.users(name, email, password) 
							VALUES($1, $2, $3)
							RETURNING id`)
	if err != nil {
		utils.HandleError(err, false)
	}

	defer stmt.Close()

	var id int64 = 0

	_, err = stmt.Exec(name, email, password)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == pgerrcode.UniqueViolation {
				utils.HandleError(utils.DupError(email), false)
				pqError := &pq.Error{Message: error.Error(utils.DupError(email))}
				
				return id, pqError
			}
		} else {
			utils.HandleError(err, false)
		}
	}

	return id, err
}