package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Assert akan memanggi fail()
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fatur")
	assert.Equal(t, "Hello Fatur", result, "Result must be Hello Fatur")
	fmt.Println("Test hello world with assert done")
}

// Require akan memanggil failNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fatu")
	require.Equal(t, "Hello Fatur", result, "Result must be Hello Fatur")
	fmt.Println("Test hello world with require done")
}