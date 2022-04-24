package database

import (
	"errors"
	"fmt"

	"github.com/DavidHODs/TechHUB-goGraph/utils"
)

// saves created posts to the database
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

func LikePostAndUpdateCount(userID, postID string) (string, error) {
	if userID == "" {
		utils.HandleError(errors.New("you have to login before you can like a post"), false)
		return "", errors.New("you have to login before you can like a post") 
	}

	stmt, err := Db.Prepare(`UPDATE tech.posts SET likes = $1 WHERE id = $2 RETURNING likes`)
	if err != nil {
		utils.HandleError(err, false)
		return "", errors.New("something went wrong, try again later")
	}

	defer stmt.Close()

	var likesID string = ""

	err = stmt.QueryRow(userID, postID).Scan(&likesID)
	if err != nil {
		utils.HandleError(err, false)
		return "", errors.New("something went wrong, try again later")
	}

	fmt.Println(likesID)

	return likesID, nil
}