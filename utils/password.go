package utils

import (
	"errors"

	"github.com/go-passwd/validator"
	"golang.org/x/crypto/bcrypt"
)

// checks for password strength, password field match
func PasswordCheck(password, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return errors.New("password fields does not match")
	}

	val := validator.New(validator.CommonPassword(errors.New("password is too common")), validator.MinLength(8, errors.New("password length must be atleast 8 characters")))

	err := val.Validate(password)
	if err != nil {
		return err
	}

	return nil
}

// hashes the password of users 
func HashPassword(password string) ([]byte, error) {
	hashedP, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		hashError := errors.New("something went wrong from our end, try again later")
		HandleError(hashError, false)
		return nil, hashError
	}

	return hashedP, nil
}

// compares if given password during login attempt matches with stored hash
func CheckHashAgainstPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}