package unit_testing

import (
	"testing"

	file_operations "github.com/RifqiWasntHere/golang-module"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNewFile(t *testing.T) {
	t.Run("SubTestCreateNewFile", func(t *testing.T) {
		result := file_operations.CreateNewFile("test.txt", "false")
		if result != nil {
			t.FailNow()
		}
	})

	t.Run("SubTestCreateNewFile2", func(t *testing.T) {
		result := file_operations.CreateNewFile("test2.txt", "false")
		if result != nil {
			t.Fail()
		}
	})
}

func BenchmarkCreateNewFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// file_operations.CreateNewFile("tweet", "Aku tweet")
		i = i + 0
	}
}

func TestReadFile(t *testing.T) {
	_, err := file_operations.ReadFile("test.txt")

	require.Equal(t, nil, err, "require == FailNow()!")
}

func TestAppendToFile(t *testing.T) {
	result := file_operations.AppendToFile("test.txt", "\n apa pula")
	assert.Equal(t, nil, result, "assert == Fail()!")
}
