package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "test@test.com", expected: true},
		{input: "t@t.c", expected: true},
		{input: "test", expected: false},
		{input: "@", expected: false},
		{input: "", expected: false},
	}

	for _, testCase := range testCases {
		got := ValidateEmail(testCase.input)
		if got != testCase.expected {
			t.Errorf("\nFor input: %s\nExpected: %t\nGot: %t", testCase.input, testCase.expected, got)
		}
	}
}
func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "12345678", expected: true},
		{input: "qwertyasdfgzxcvb", expected: true},
		{input: "1234567", expected: false},
		{input: "wasd", expected: false},
		{input: "", expected: false},
	}

	for _, testCase := range testCases {
		got := ValidatePassword(testCase.input)
		if got != testCase.expected {
			t.Errorf("\nFor input: %s\nExpected: %t\nGot: %t", testCase.input, testCase.expected, got)
		}
	}
}
