package unit_testing

import (
	"testing"

	file_operations "github.com/RifqiWasntHere/golang-module"
)

func TestCreateNewFile(t *testing.T) {
	result := file_operations.CreateNewFile("test.txt", "false")
	if result != nil {
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	_, err := file_operations.ReadFile("test.txt")
	if err != nil {
		t.Fatal("ada yang salah ini")
	}
}

func TestAppendToFile(t *testing.T) {
	result := file_operations.AppendToFile("tet.txt", "\n apa pula")
	if result != nil {
		t.Fatal("ada yang salah ini")
	}
}
