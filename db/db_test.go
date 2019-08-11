package db

import (
	"os"
	"testing"
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
		path := getDBPath()
		if path != testCase.expected {
			t.Errorf("\nFor input %s\nGot: %s\nExpected: %s", testCase.input, path, testCase.expected)
		}
	}
}
