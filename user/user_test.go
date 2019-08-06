package user

import (
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	testCases := []struct {
		test     User
		expected error
	}{
		{
			test:     User{Password: "12345678"},
			expected: nil,
		},
		{
			test:     User{Password: "qwertyasdfzxcv"},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		got := testCase.test.EncryptPassword()
		if got != testCase.expected {
			t.Errorf("\nFor input: %v\nExpected: %v\nGot: %v", testCase.test, testCase.expected, got)
		}
	}
}

func TestValid(t *testing.T) {
	testCases := []struct {
		test     User
		expected bool
	}{
		{
			test:     User{Name: "test", Email: "test@test.com", Password: "12345678"},
			expected: true,
		},
		{
			test:     User{Name: "", Email: "test@test.com", Password: "12345678"},
			expected: false,
		},
		{
			test:     User{Name: "test", Email: "", Password: "12345678"},
			expected: false,
		},
		{
			test:     User{Name: "test", Email: "test@test.com", Password: ""},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		got := testCase.test.Valid()
		if got != testCase.expected {
			t.Errorf("\nFor input: %v\nExpected: %t\nGot: %t", testCase.test, testCase.expected, got)
		}
	}
}