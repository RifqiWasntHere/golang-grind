package unit_testing

import (
	"testing"

	file_operations "github.com/RifqiWasntHere/golang-module"
)

func TestTableCreateNewFile(t *testing.T) {
	sequences := []struct {
		testName string
		fileName string
		content  string
	}{
		{
			testName: "TestTableCreateNewFile1",
			fileName: "table1",
			content:  "this is table 1",
		},
		{
			testName: "TestTableCreateNewFile2",
			fileName: "table2",
			content:  "this is table 2",
		},
		{
			testName: "TestTableCreateNewFile3",
			fileName: "table3",
			content:  "this is table 3",
		},
	}

	for _, sequence := range sequences {
		t.Run(sequence.testName, func(t *testing.T) {
			result := file_operations.CreateNewFile(sequence.fileName, sequence.content)
			if result != nil {
				t.FailNow()
			}
		})
	}
}
