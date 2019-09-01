package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBPath(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		panic("Couldn't get the current working directory")
	}

	testCases := []struct {
		input    string
		expected string
	}{
		{input: "TEST", expected: currentPath + dbTestPath},
		{input: "WHATEVER", expected: currentPath + dbPath},
	}

	for _, testCase := range testCases {
		os.Setenv("ENV", testCase.input)
		got := getDBPath()
		assert.Equal(t, testCase.expected, got)
	}
}
