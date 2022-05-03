package database

import (
	"database/sql"
	"fmt"

	"github.com/DavidHODs/TechHUB-goGraph/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Db *sql.DB

// creates a persistent connection to the database and migrates latest schema changes 
func ConnectAndMigrate() {
	_, host, dbport, user, password, dbname, sslmode, _ := utils.LoadEnv()
	
	databaseInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, dbport, user, password, dbname, sslmode)
	
	db, err := sql.Open("postgres", databaseInfo)
	if err != nil {
		utils.HandleError(err, true)
	}

	err = db.Ping()
	if err != nil {
		utils.HandleError(err, true)
	}

	fmt.Println("database connection successful")

	Db = db

	err = Db.Ping()
	if err != nil {
		utils.HandleError(err, true)
	}

	driver, _ := postgres.WithInstance(Db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://postgres/migrations/postgres",
		fmt.Sprintf("postgres://%s:%s@/%s", user, password, dbname),
		driver,
		
	)

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		utils.HandleError(err, true)
	}

	fmt.Println("database migration successful")
}