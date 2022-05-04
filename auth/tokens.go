package auth

import (
	"context"
	"errors"
	"time"

	"github.com/DavidHODs/TechHUB-goGraph/utils"
	"github.com/dgrijalva/jwt-go"
)

// Generates a token and assigns an email to its claims
func GenerateToken(ctx context.Context, email string) (string, error) {
	_, _, _, _, _, _, _, key := utils.LoadEnv()
	secretKey := []byte(key)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	// claims to the token expires after 24 hrs 
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		utils.HandleError(errors.New("error in generating key"), false)
		return "", err
	}
	
	_ = context.WithValue(ctx, "AuthToken", tokenString)
	return tokenString, nil
}

// parses a token and returns the email in its claims 
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