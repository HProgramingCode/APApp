package testutils

import (
	"main/internal/middleware"
	"os"
	"testing"

	"go.uber.org/zap"
)

// SetupTestEnv initializes shared test environment (logger, etc.)
func SetupTestEnv() {
	middleware.Log, _ = zap.NewDevelopment()
}

// RunTests runs all tests after setup
func RunTests(m *testing.M) {
	SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}
