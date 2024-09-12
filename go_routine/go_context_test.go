package go_routine

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextB, "c", "C")

	//Context's able to fetch it's own value, and parent's value. but unable to fetch child's values
	fmt.Println(contextB)
	fmt.Println(contextB.Value("b"))
	fmt.Println(contextB.Value("c"))

	fmt.Println(contextC)
	fmt.Println(contextC.Value("c"))
	fmt.Println(contextC.Value("b"))
}
