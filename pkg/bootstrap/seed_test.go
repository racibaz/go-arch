package bootstrap

import (
	"testing"
)

func TestSeed_NoPanic(t *testing.T) {
	// This test will only check that Seed does not panic.
	// In real scenarios, you should mock config, database, and seeder sub-calls!
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Seed panicked: %v", r)
		}
	}()
	Seed()
}
