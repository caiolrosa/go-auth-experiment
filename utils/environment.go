package utils

import "os"

const (
	// TestEnv is the test environment variable key
	TestEnv = "TEST"
)

// GetEnv returns the environment runtime
func GetEnv() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		return TestEnv
	}

	return env
}
