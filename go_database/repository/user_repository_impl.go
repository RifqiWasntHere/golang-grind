package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go_database/model"
	"log"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func Nullable(col sql.NullString) string { //Jorking of null columns
	if col.Valid {
		return col.String
	}
	return ""
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user model.User) (model.User, error) {

	stmt, err := repository.DB.PrepareContext(ctx, "INSERT INTO Users(userId, userName) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, user.UserId, user.UserName, user.UserOccupation).Scan(&user.UserName, &user.UserOccupation)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, user model.User) (model.User, error) {
	query := "SELECT userId, userName, userOccupation FROM Users WHERE userId = ?"

	// Execute the query
	rows, err := repository.DB.QueryContext(ctx, query, user.UserId)
	if err != nil {
		log.Fatal("Query failed: %v", err)
	}
	defer func() {
		if rows != nil {
			rows.Close() // Close rows if they were opened successfully
		}
	}()

	// Iterate over the rows
	for rows.Next() {
		var userId, userName string
		var userOccupation sql.NullString

		// Scan the row into the variables
		err = rows.Scan(&userId, &userName, &userOccupation)
		if err != nil {
			log.Fatal("Row scan failed: %v", err)
		}

		userOccupationCol := Nullable(userOccupation)
		fmt.Printf("User ID: %s\nUsername: %s\nOccupation: %s\n", userId, userName, userOccupationCol)
	}
}
