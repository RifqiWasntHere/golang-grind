package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	payload := "rifqi"

	//Encode
	encode := base64.StdEncoding.EncodeToString([]byte(payload))
	fmt.Println(encode)

	//Decode
	decode, err := base64.StdEncoding.DecodeString(encode) //Function ini diassign ke 2 variable karena functionnya men-return byte dan error

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decode))
	}
}
