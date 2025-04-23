package controller

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Returns a formatted mysql connection string
func loadCredentials() string {
	err := godotenv.Load("creds.env")
	if err != nil {
		log.Fatalln("Error loading credentials file: ", err.Error())
	}

	//user:pass@tcp(server:port)/dbname

	var (
		server = os.Getenv("DB_SERVER")
		port   = os.Getenv("DB_PORT")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		dbname = os.Getenv("DB_NAME")
	)

	dbstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, server, port, dbname)

	return dbstring
}

// Returns a pointer to the database connection so we can make requests to the database
func NewConnection() *sql.DB {
	var DB *sql.DB
	var err error
	DB, err = sql.Open("mysql", loadCredentials())
	if err != nil {
		log.Fatalln("Couldn't establish the connection with the database due to the following error: ", err.Error())
	}

	return DB
}
