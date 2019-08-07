package db

import ( 
	"testing"
	"os"
)

func TestGetDBPath(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		panic("Couldn't get the current working directory")
	}

	testCases := []struct {
		input string
		expected string
	} {
		{input: "TEST", expected: currentPath+gormTestDBPath},
		{input: "WHATEVER", expected: currentPath+gormDBPath},
	}

	for _, testCase := range testCases {
		os.Setenv("ENV", testCase.input)
		path := getDBPath()
		if path != testCase.expected {
			t.Errorf("\nFor input %s\nGot: %s\nExpected: %s", testCase.input, path, testCase.expected)
		}
	}
}
