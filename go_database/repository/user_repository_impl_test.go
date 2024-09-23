package repository

import (
	"context"
	"fmt"
	"go_database/db"
	"go_database/model"
	"testing"
)

func TestUserInsert(t *testing.T) {
	userRepository := NewUserRepository(db.GetDatabase())

	ctx := context.Background()
	user := model.User{
		UserId:         "76123761",
		UserName:       "Niggadongdong",
		UserOccupation: "stripper",
	}

	result, err := userRepository.Insert(ctx, user)
	if err != nil {
		t.Fatal("error! :", err)
	}

	fmt.Println(result)
}

func TestUserFindById(t *testing.T) {
	userRepository := NewUserRepository(db.GetDatabase())

	ctx := context.Background()

	result, err := userRepository.FindById(ctx, "12312")
	if err != nil {
		t.Fatal("error ! :", err)
	}

	fmt.Println(result)
}
