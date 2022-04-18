package utils

import (
	"errors"
	"log"
	"os"
	"runtime"
)

// HandleError is a custom error function that prints to the terminal an error message, the function throwing the error, the filename and code line number responsible for the error. It takes two parameter: the error message and a boolean condition. When set to true, it replicates log.Fatal()
func HandleError(err error, fatal bool) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s :: [%s :line %d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
	}

	if fatal {
		os.Exit(1)
	}
}

// returns a duplicate error message which includes the duplicated field
func DupError(err string) error {
	return errors.New(err + DuplicateError)
}

// A key value pair for error type and returned error messages
const (
	DuplicateError = " record already exist"
)