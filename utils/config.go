package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// loads the environment variables
func LoadEnv() (string, string, int, string, string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		HandleError(err, true)
	}

	dbport, _ := strconv.Atoi(os.Getenv("dbport"))
	port := os.Getenv("port")
	host :=  os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	sslmode := os.Getenv("sslmode")

	return port, host, dbport, user, password, dbname, sslmode
}