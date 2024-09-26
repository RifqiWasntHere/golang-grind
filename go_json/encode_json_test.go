package go_json

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func logJson(payload interface{}) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Marshal error : ", err)
	}

	fmt.Println(string(bytes))
}
func TestEncodeJson(t *testing.T) {
	logJson("Hey There")
	logJson(1)
	logJson(true)
	logJson([]string{"Muhamad", "Rifqi", "Fadhillah"})
}
