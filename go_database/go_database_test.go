package go_database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TestOpenConnection(t *testing.T) {
	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env")
	}

	// Retrieve environment variables
	dbCreds := os.Getenv("DB_CREDS")
	dbName := os.Getenv("DB_NAME")

	// open db connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s", dbCreds, dbName))
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	db.SetMaxIdleConns(5)                   //Max. amount of idle dbconn
	db.SetMaxOpenConns(10)                  //Max. amount of dbconn at a same time
	db.SetConnMaxIdleTime(5 * time.Minute)  //Max. lifetime a dbconn can idle
	db.SetConnMaxLifetime(60 * time.Minute) //Max. lifetime of a dbconn

	// defer db.Close() //Can't use Close() when setting up orjlidle dbconns

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		fmt.Println("Connected to the database successfully!")
	}
}
