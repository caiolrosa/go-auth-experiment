package utils

import (
	"os"
	"testing"
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
		env := GetEnv()
		if env != testCase.expected {
			t.Errorf("\nFor input: %s\nExpected: %s\nGot: %s", testCase.input, testCase.expected, env)
		}
	}
}
