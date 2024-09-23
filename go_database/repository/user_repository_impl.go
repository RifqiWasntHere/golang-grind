package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_database/model"
	"log"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func Nullable(col sql.NullString) string { //Jorking of null columns
	if col.Valid {
		return col.String
	}
	return ""
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user model.User) (model.User, error) {

	stmt, err := repository.DB.PrepareContext(ctx, "INSERT INTO Users(userId, userName, userOccupation) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, user.UserId, user.UserName, user.UserOccupation).Scan(&user.UserId, &user.UserName, &user.UserOccupation)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, id string) (model.User, error) {
	query := "SELECT userId, userName, userOccupation FROM Users WHERE userId = ?"
	user := model.User{}
	// Execute the query
	rows, err := repository.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal("Query failed: ", err)
	}
	defer rows.Close()

	// Iterate over the rows
	if rows.Next() {
		// Scan the row into the variablesc
		rows.Scan(&user.UserId, &user.UserName, &user.UserOccupation)
		return user, nil
	} else {
		return user, errors.New("User id : " + id + " is not found")
	}
}
