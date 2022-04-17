package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/DavidHODs/TechHUB-goGraph/customerror"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func LoadEnv() (string, string, int, string, string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		customerror.HandleError(err, true)
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

func ConnectAndMigrate() {
	_, host, dbport, user, password, dbname, sslmode := LoadEnv()
	
	databaseInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, dbport, user, password, dbname, sslmode)
	
	db, err := sql.Open("postgres", databaseInfo)
	if err != nil {
		customerror.HandleError(err, true)
	}

	err = db.Ping()
	if err != nil {
		customerror.HandleError(err, true)
	}

	fmt.Println("database connection successful")

	Db = db

	err = Db.Ping()
	if err != nil {
		customerror.HandleError(err, true)
	}

	driver, _ := postgres.WithInstance(Db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://postgres/migrations/postgres",
		fmt.Sprintf("postgres://%s:%s@/%s", user, password, dbname),
		driver,
		
	)

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		customerror.HandleError(err, true)
	}

	fmt.Println("database migration successful")
}