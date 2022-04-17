package customerror

import (
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