package auth

import (
	"database/sql"
	"errors"

	myDB "github.com/DavidHODs/TechHUB-goGraph/postgres"
	"github.com/DavidHODs/TechHUB-goGraph/utils"
)

func Authenticate(email, password string) bool {
	stmt, err := myDB.Db.Prepare(`SELECT password from tech.users WHERE email = $1`)
	if err != nil {
		utils.HandleError(err, false)
	}

	defer stmt.Close()

	var hashedPassword string = ""

	err = stmt.QueryRow(email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.HandleError(errors.New("user does not exist"), false)
			return false
		}
	}

	return utils.CheckHashAgainstPassword(hashedPassword, password)
}