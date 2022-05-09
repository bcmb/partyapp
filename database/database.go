package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var DbConn *sql.DB

var migrationsDirectory = filepath.Join("database", "migrations")

var user = os.Getenv("DB_PASSWORD")
var password = os.Getenv("DB_USER")
var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var db = os.Getenv("DB_NAME")

func SetupDb(shouldMigrate bool) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
	var err error
	DbConn, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal("Failed to parse database connection params %s", err)
	}

	err = DbConn.Ping()

	if err != nil {
		log.Fatal("Failed to establish connection to database %s", err)
	} else {
		if shouldMigrate {
			runMigrations()
		}
	}
}

func runMigrations() {
	scripts, err := ioutil.ReadDir(migrationsDirectory)

	if err != nil {
		log.Fatal("Failed to read migrations folder contents %s", err)
	}

	for _, file := range scripts {
		content, err := ioutil.ReadFile(filepath.Join(migrationsDirectory, file.Name()))

		if err != nil {
			log.Fatalf("Failed to read file %s %s", file.Name(), err)
		}

		_, err = DbConn.Exec(string(content))

		if err != nil {
			log.Fatalf("Failed to execute migration script %s %s", file.Name(), err)
		}
	}

}
