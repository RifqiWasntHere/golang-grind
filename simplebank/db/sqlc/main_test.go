package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/util"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..") // ../.. means "go to parent folder"
	if err != nil {
		log.Fatal("can't load configuration files")
	}
	// testDb, err := sql.Open(dbDriver, dbSource) // Ini kode lama. ini buat error karena instead of assign sql.Open ke global variabel, itu malah bikin testDb baru di scope function
	testDb, err = sql.Open(config.DBDriver, config.DBSource) // Use "=" to assign to the global testDb
	if err != nil {
		log.Fatal("can't connect to the database:", err)
	}

	defer testDb.Close() // Close the database connection after all tests have run

	testQueries = New(testDb)

	// Run the tests
	os.Exit(m.Run())
}
