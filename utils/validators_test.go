package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		test     string
		expected bool
	}{
		{test: "test@test.com", expected: true},
		{test: "t@t.c", expected: true},
		{test: "test", expected: false},
		{test: "@", expected: false},
		{test: "", expected: false},
	}

	for _, testCase := range testCases {
		got := ValidateEmail(testCase.test)
		if got != testCase.expected {
			t.Errorf("\nFor input: %s\nExpected: %t\nGot: %t", testCase.test, testCase.expected, got)
		}
	}
}
func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		test     string
		expected bool
	}{
		{test: "12345678", expected: true},
		{test: "qwertyasdfgzxcvb", expected: true},
		{test: "1234567", expected: false},
		{test: "wasd", expected: false},
		{test: "", expected: false},
	}

	for _, testCase := range testCases {
		got := ValidatePassword(testCase.test)
		if got != testCase.expected {
			t.Errorf("\nFor input: %s\nExpected: %t\nGot: %t", testCase.test, testCase.expected, got)
		}
	}
}
