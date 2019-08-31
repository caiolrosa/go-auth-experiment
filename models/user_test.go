package models

import (
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	testCases := []struct {
		input    User
		expected error
	}{
		{input: User{Password: "12345678"}, expected: nil},
		{input: User{Password: "qwertyasdfzxcv"}, expected: nil},
	}

	for _, testCase := range testCases {
		got := testCase.input.EncryptPassword()
		if got != testCase.expected {
			t.Errorf("\nFor input: %v\nExpected: %v\nGot: %v", testCase.input, testCase.expected, got)
		}
	}
}

func TestAuthenticate(t *testing.T) {
	testCases := []struct {
		base     User
		input    string
		expected error
	}{
		{base: User{Password: "12345678"}, input: "12345678", expected: nil},
		{base: User{Password: "abcdef"}, input: "abcdef", expected: nil},
	}

	for _, testCase := range testCases {
		if err := testCase.base.EncryptPassword(); err != nil {
			t.Error(err)
			continue
		}

		if got := testCase.base.Authenticate(testCase.input); got != testCase.expected {
			t.Errorf("\nFor base: %v and the input: %s\nExpected: %s\nGot: %s",
				testCase.base, testCase.input, testCase.expected.Error(), got.Error())
		}
	}

	failUser := User{Password: "passwordtofail"}
	if err := failUser.EncryptPassword(); err != nil {
		t.Error(err)
		return
	}

	if got := failUser.Authenticate("incorrectpassword"); got == nil {
		t.Errorf("\nFor base: %v and the input: %s\nExpected: %s\nGot: %s",
			failUser, "incorrectpassword", "An error", got.Error())
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
