package utils

import (
	"errors"

	"github.com/go-passwd/validator"
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