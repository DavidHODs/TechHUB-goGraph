package auth

import (
	"errors"
	"time"

	"github.com/DavidHODs/TechHUB-goGraph/utils"
	"github.com/dgrijalva/jwt-go"
)


func GenerateToken(email string) (string, error) {
	_, _, _, _, _, _, _, key := utils.LoadEnv()
	secretKey := []byte(key)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		utils.HandleError(errors.New("error in generating key"), false)
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	_, _, _, _, _, _, _, key := utils.LoadEnv()
	secretKey := []byte(key)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	} else {
		return "", err
	}
}