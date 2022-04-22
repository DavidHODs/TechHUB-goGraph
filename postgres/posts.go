package database

import (
	"errors"

	"github.com/DavidHODs/TechHUB-goGraph/utils"
)


func SavePost(author, body, sharedBody, image string) (string, string, error) {
	if body == "" {
		utils.HandleError(errors.New("post can not be blank"), false)
		return "", "", errors.New("post can not be blank")
	}

	if author == "" {
		utils.HandleError(errors.New("author can not be blank"), false)
		return "", "", errors.New("author can not be blank")
	}

	postLength := len(body)
	if postLength > 1024 {
		utils.HandleError(errors.New("post can not be longer than 1024 words"), false)
		return "", "", errors.New("post can not be longer than 1024 words")
	}

	stmt, err := Db.Prepare(`INSERT INTO tech.posts(author, body, shared_body, image)
							VALUES($1, $2, $3, $4)
							RETURNING id`)
	if err != nil {
		utils.HandleError(err, false)
		return "", "", errors.New("something went wrong, try reposting")
	}

	defer stmt.Close()

	var id string = ""

	err = stmt.QueryRow(author, body, sharedBody, image).Scan(&id)
	if err != nil {
		utils.HandleError(err, false)
		return "", "", errors.New("something went wrong, try reposting")
	}

	return id, body, nil
}