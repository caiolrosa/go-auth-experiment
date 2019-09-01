package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, testCase.expected, got)
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

		got := testCase.base.Authenticate(testCase.input)
		assert.Equal(t, testCase.expected, got)
	}

	failUser := User{Password: "password_to_fail"}
	if err := failUser.EncryptPassword(); err != nil {
		t.Error(err)
		return
	}

	got := failUser.Authenticate("incorrect_password")
	assert.NotEqual(t, got, nil)
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
		assert.Equal(t, testCase.expected, got)
	}
}
