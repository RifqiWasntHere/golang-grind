package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	// testDb, err := sql.Open(dbDriver, dbSource) // Ini kode lama. ini buat error karena instead of assign sql.Open ke global variabel, itu malah bikin testDb baru di scope function
	testDb, err = sql.Open(dbDriver, dbSource) // Use "=" to assign to the global testDb
	if err != nil {
		log.Fatal("can't connect to the database:", err)
	}

	defer testDb.Close() // Close the database connection after all tests have run

	testQueries = New(testDb)

	// Run the tests
	os.Exit(m.Run())
}
