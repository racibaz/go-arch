package trace

import (
	"testing"
)

func TestInitTracer(t *testing.T) {
	tp, err := InitTracer()
	// The exporter requires a running Jaeger endpoint at jaeger:4318 (or will error)
	if err == nil && tp == nil {
		t.Error("Expected non-nil TracerProvider if no error returned")
	}
	if err != nil {
		t.Logf("InitTracer returned error (expected if Jaeger is missing): %v", err)
	} else {
		t.Logf("InitTracer succeeded, TracerProvider: %v", tp)
	}
}
