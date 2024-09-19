package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO Users(userId, userName) VALUES('21132123', 'Rifqi')")
	if err != nil {
		fmt.Println(err)
	}
}

func TestInject(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()

	userName := "Liu Bang"
	married := true

	query := "SELECT userName, occupation FROM UsersComplex WHERE userName = ? AND married = ?"
	rows, err := db.QueryContext(ctx, query, userName, married)

	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userName, occupation string

		err = rows.Scan(&userName, &occupation)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Username : %s\nOccupation : %s\n", userName, occupation)
	}

}

func TestPrepareInject(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO Users(userId, userName) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	id, name := "1231212", "koproy"
	err = stmt.QueryRowContext(ctx, id, name).Scan(&id, &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data Inserted")
}

// Basic Table
func TestRead(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT userId, userName, userOccupation FROM Users"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		fmt.Println("Query Failed")
	}

	defer rows.Close()

	for rows.Next() {
		var userId, userName string
		var userOccupation sql.NullString

		err = rows.Scan(&userId, &userName, &userOccupation)
		if err != nil {
			fmt.Println(err)
		}

		userOccupationCol := Nullable(userOccupation)
		fmt.Printf("User ID : %s\nUsername : %s\nOccupation : %s\n", userId, userName, userOccupationCol)
	}
}

func Nullable(col sql.NullString) string { //Jorking of null columns
	if col.Valid {
		return col.String
	}
	return ""
}

// Complex Table
func TestInsertComplex(t *testing.T) {

}

func TestReadComplex(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()

	// query := "SELECT userId, userName, email, balance, score, birthdate, married, occupation, createdAt FROM UsersComplex"
	query := "SELECT * FROM UsersComplex"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		fmt.Println("Query Failed")
	}

	defer rows.Close()

	for rows.Next() {
		var userId, userName string
		var occupation, email sql.NullString
		var balance int32
		var score float32
		var birthdate, createdAt time.Time
		var married bool

		err = rows.Scan(&userId, &userName, &email, &balance, &score, &birthdate, &married, &occupation, &createdAt)
		if err != nil {
			fmt.Println(err)
		}

		occupationCol := Nullable(occupation)
		emailCol := Nullable(email)
		fmt.Printf("User ID : %s\nUsername : %s\nOccupation : %s\nemail : %s\n", userId, userName, occupationCol, emailCol)
		fmt.Println("Balance :", balance)
		fmt.Println("Score :", score)
		fmt.Println("Birthdate :", birthdate)
		fmt.Println("married :", married)
		fmt.Println("Created At :", createdAt)
	}
}
