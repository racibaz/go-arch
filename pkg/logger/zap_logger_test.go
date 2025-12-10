package logger

import (
	"testing"
)

func TestNewZapLogger(t *testing.T) {
	zl, err := NewZapLogger()
	if err != nil {
		t.Errorf("NewZapLogger returned error: %v", err)
	}
	if zl == nil {
		t.Fatalf("Expected non-nil ZapLogger")
	}
}

func TestZapLoggerMethods(t *testing.T) {
	zl, err := NewZapLogger()
	if err != nil {
		t.Fatalf("Failed to create ZapLogger: %v", err)
	}
	zl.Debug("debug message: %s", "test")
	zl.Info("info message: %s", "test")
	zl.Warn("warn message: %s", "test")
	zl.Error("error message: %s", "test")
	// Avoid calling Fatal directly as it will exit the test process.
}
