package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Skills struct {
	Name        string
	Proficiency string
}
type CustomerComplex struct {
	Name    string
	Email   string
	Hobbies []string // Regular slice
	Skills  []Skills // Slice of struct
}

func TestJsonArray(t *testing.T) {

	customer := CustomerComplex{
		Name:    "Rifqi",
		Email:   "Rifqi@Hotmail.com",
		Hobbies: []string{"Badminton", "Mewing", "Jelqing"},
		Skills: []Skills{
			{
				Name:        "Golang",
				Proficiency: "Basic",
			},
			{
				Name:        "NodeJs",
				Proficiency: "Intermediate",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

// Directly Decode Json Array of Objects Into a Slice
func TestDecodeJsonArray(t *testing.T) {

	customer := `[{"Name":"Golang","Proficiency":"Basic"},{"Name":"NodeJs","Proficiency":"Intermediate"}]`

	payload := &[]Skills{}
	err := json.Unmarshal([]byte(customer), payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(payload)
}
