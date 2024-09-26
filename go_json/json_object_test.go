package go_json

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Customer struct {
	Name     string
	Email    string
	IsActive bool
}

func TestJsonObject(t *testing.T) {
	payload := Customer{
		Name:     "Rifqi",
		Email:    "RifqiF@gmail.com",
		IsActive: true,
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Marshal error : ", err)
	}

	fmt.Println(string(bytes))
}

func TestJsonDecode(t *testing.T) {
	jsonString := `{"Name":"Rifqi","Email":"RifqiF@gmail.com","IsActive":true}`
	jsonBytes := []byte(jsonString)

	payload := &Customer{}

	err := json.Unmarshal(jsonBytes, payload)
	if err != nil {
		log.Fatal("Unmarshal error :", err)
	}
	fmt.Println(payload)
}
