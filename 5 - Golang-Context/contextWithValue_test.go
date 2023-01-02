package go_context

import (
	"context"
	"fmt"
	"testing"
)

func TestCtxWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Get value child
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
}