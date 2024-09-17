package go_database

import (
	"context"
	"database/sql"
	"fmt"
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
