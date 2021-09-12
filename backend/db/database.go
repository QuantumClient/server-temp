package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Db *sql.DB

func Init() {
	godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbURL := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	if db, err := sql.Open("mysql", dbURL); err != nil {
		log.Print(err.Error())
	} else {
		Db = db
	}

}
