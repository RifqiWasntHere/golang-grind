package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError = errors.New("Not Found Error")
)

func getUserById(userid string) error {
	if userid ==""{
		return ValidationError
	}

	if userid!= "Rifi"{
		return NotFoundError
	}

	return nil
}

func main() {
	result := getUserById("Rifqi")
	if result != nil {
		if errors.Is(result, ValidationError){
			fmt.Println(result)
		} else if errors.Is(result, NotFoundError){
			fmt.Println(result)
		} else {
			fmt.Println("Error not found ini mah")
		}
	}
}