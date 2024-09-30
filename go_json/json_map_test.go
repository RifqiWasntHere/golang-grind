package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMap(t *testing.T) {
	jsonPayload := `{"name": "Rifqi", "user_occupation":"Professional Yapper"}`

	var jsonMap map[string]interface{}
	_ = json.Unmarshal([]byte(jsonPayload), &jsonMap)

	fmt.Println(jsonMap)
	fmt.Println(jsonMap["user_occupation"]) // Just Index The JSON Key !
}

func TestJsonMapEncode(t *testing.T) {
	jsonMap := map[string]interface{}{
		"name":            "Rifqi",
		"user_occupation": "Professional Yapper",
	}

	bytes, _ := json.Marshal(jsonMap)

	println(string(bytes))
}
