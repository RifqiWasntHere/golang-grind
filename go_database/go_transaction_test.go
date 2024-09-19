package go_database

import (
	"context"
	"fmt"
	"go_database/db"
	"log"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := db.GetDatabase()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Sumtings Wong wih deh TX")
	}

	query := "INSERT INTO Users(userId, userName) VALUES(?, ?)"

	id, name := "561212", "cihuy"
	result, err := tx.ExecContext(ctx, query, id, name)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit() //Commit to execute queries
	// err = tx.Rollback() //Rollback to abort queries
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
