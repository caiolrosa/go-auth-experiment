package user

import (
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	testCases := []struct {
		input    User
		expected error
	}{
		{
			input:    User{Password: "12345678"},
			expected: nil,
		},
		{
			input:    User{Password: "qwertyasdfzxcv"},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		got := testCase.input.EncryptPassword()
		if got != testCase.expected {
			t.Errorf("\nFor input: %v\nExpected: %v\nGot: %v", testCase.input, testCase.expected, got)
		}
	}
}

func TestValid(t *testing.T) {
	testCases := []struct {
		input    User
		expected bool
	}{
		{
			input:    User{Name: "test", Email: "test@test.com", Password: "12345678"},
			expected: true,
		},
		{
			input:    User{Name: "", Email: "test@test.com", Password: "12345678"},
			expected: false,
		},
		{
			input:    User{Name: "test", Email: "", Password: "12345678"},
			expected: false,
		},
		{
			input:    User{Name: "test", Email: "test@test.com", Password: ""},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		got := testCase.input.Valid()
		if got != testCase.expected {
			t.Errorf("\nFor input: %v\nExpected: %t\nGot: %t", testCase.input, testCase.expected, got)
		}
	}
}
