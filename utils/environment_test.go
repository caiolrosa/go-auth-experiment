package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "TEST", expected: "TEST"},
		{input: "PROD", expected: "PROD"},
		{input: "", expected: "TEST"},
	}

	for _, testCase := range testCases {
		os.Setenv("ENV", testCase.input)
		got := GetEnv()
		assert.Equal(t, testCase.expected, got)
	}
}
