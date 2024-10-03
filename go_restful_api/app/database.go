package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseCreds struct {
	Credential   string
	DatabaseName string
}

func NewDB(dbCreds *DatabaseCreds) *sql.DB {

	// open db connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s?parseTime=true", dbCreds.Credential, dbCreds.DatabaseName))
	if err != nil {
		log.Fatalln("Failed to connect to database", err)
	}

	db.SetMaxIdleConns(5)                   //Max. amount of idle dbconn
	db.SetMaxOpenConns(10)                  //Max. amount of dbconn at a same time
	db.SetConnMaxIdleTime(5 * time.Minute)  //Max. lifetime a dbconn can idle
	db.SetConnMaxLifetime(60 * time.Minute) //Max. lifetime of a dbconn

	return db
}
