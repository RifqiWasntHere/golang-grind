package main

import (
	"fmt"

	file_operations "github.com/RifqiWasntHere/golang-module"
)

func main() {
	file_operations.CreateNewFile("sigma.txt", "Aku real smigma")
	file_operations.AppendToFile("sigma.txt", " dan Kamu ligma ")
	result, _ := file_operations.ReadFile("sigma.txt")
	fmt.Println(result)
}
