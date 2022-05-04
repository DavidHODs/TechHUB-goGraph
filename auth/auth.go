package auth

import (
	myDB "github.com/DavidHODs/TechHUB-goGraph/postgres"
	"github.com/DavidHODs/TechHUB-goGraph/utils"
)

// pulls up the stored password hash via user email and checks if the supplied password on login attempt matches
func Authenticate(email, password string) bool {
	_, hashedPassword, _, err := myDB.GetUserDetailsByEmail(email)
	if err != nil {
		utils.HandleError(err, false)
		return false
	}

	return utils.CheckHashAgainstPassword(hashedPassword, password)
}