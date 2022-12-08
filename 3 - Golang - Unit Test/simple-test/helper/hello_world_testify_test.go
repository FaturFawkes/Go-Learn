package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Main Testing
func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")
	
	m.Run()

	//after
	fmt.Println("After Unit Test")
}

// Assert akan memanggi fail()
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fatur")
	assert.Equal(t, "Hello Fatur", result, "Result must be Hello Fatur")
	fmt.Println("Test hello world with assert done")
}

// Require akan memanggil failNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fatur")
	require.Equal(t, "Hello Fatur", result, "Result must be Hello Fatur")
	fmt.Println("Test hello world with require done")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("OS windows tidak bisa testing")
	}
	result := HelloWorld("Fatur")
	require.Equal(t, "Hello Fatur", result, "Result must be Hello Fatur")
}

// Sub Test
func TestHelloSubTest(t *testing.T) {
	t.Run("Fatur", func(t *testing.T) {
		result := HelloWorld("Fatur")
		require.Equal(t, "Hello Fatur", result, "Result must be Hello Fatur")
	})
	t.Run("Rohmam", func(t *testing.T) {
		result := HelloWorld("Rohman")
		require.Equal(t, "Hello Rohman", result, "Result must be Hello Fatur")
	})
}

// Table Test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Fatur",
			request: "Fatur",
			expected: "Hello Fatur",
		},
		{
			name: "Rohman",
			request: "Rohman",
			expected: "Hello Rohman",
		},
	}
	
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result not equal")
		})
	}
}