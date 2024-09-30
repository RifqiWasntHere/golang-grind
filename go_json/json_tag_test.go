package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
The core of the problem is that, GO's best practice is to name Struct Fields in PascalCase
But JSON has best practice to name Key Fields in snake_case
And for that, GO's json package supports JSON tagging ;)
*/
type CustomerTagged struct {
	Name           string `json:"name"`
	UserOccupation string `json:"user_occupation"`
}

func TestJsonTag(t *testing.T) {
	customer := CustomerTagged{
		Name:           "Rifqi",
		UserOccupation: "Pemain Clash Lane",
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	customer := `{"name":"Rifqi","user_occupation":"Pemain Clash Lane"}`

	payload := &CustomerTagged{}
	_ = json.Unmarshal([]byte(customer), payload)

	fmt.Println(payload)
}
