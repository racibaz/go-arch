package uuid_test

import (
	"testing"

	"github.com/google/uuid"
	myuuid "github.com/racibaz/go-arch/pkg/uuid"
)

func TestNewUuid_ShouldGenerateValidUUID(t *testing.T) {
	u := myuuid.NewUuid()
	if u == nil {
		t.Fatal("expected non-nil Uuid")
	}
	if u.Uuid == nil {
		t.Fatal("expected non-nil inner uuid")
	}
	if u.ToString() == "" {
		t.Fatal("expected non-empty uuid string")
	}

	_, err := uuid.Parse(u.ToString())
	if err != nil {
		t.Fatalf("invalid uuid string: %v", err)
	}
}

func TestToString_ShouldReturnEmpty_WhenNilReceiver(t *testing.T) {
	var u *myuuid.Uuid
	if u.ToString() != "" {
		t.Fatal("expected empty string for nil receiver")
	}
}

func TestToString_ShouldReturnEmpty_WhenInnerUUIDIsNil(t *testing.T) {
	u := &myuuid.Uuid{}
	if u.ToString() != "" {
		t.Fatal("expected empty string when inner uuid is nil")
	}
}

func TestParse_ShouldParseValidUUID(t *testing.T) {
	id := uuid.New().String()

	parsed, err := myuuid.Parse(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.Uuid == nil {
		t.Fatal("expected non-nil uuid")
	}
	if parsed.ToString() != id {
		t.Fatalf("expected %s, got %s", id, parsed.ToString())
	}
}

func TestParse_ShouldReturnError_OnInvalidUUID(t *testing.T) {
	_, err := myuuid.Parse("not-a-uuid")
	if err == nil {
		t.Fatal("expected error for invalid uuid")
	}
}

func TestParse_ShouldReturnEmpty_OnNilUUIDString(t *testing.T) {
	parsed, err := myuuid.Parse(uuid.Nil.String())
	if err != nil {
		t.Fatalf("did not expect error: %v", err)
	}
	if parsed.Uuid != nil {
		t.Fatal("expected nil uuid for Nil UUID")
	}
}

func TestNewID_ShouldReturnValidUUIDString(t *testing.T) {
	id := myuuid.NewID()
	if id == "" {
		t.Fatal("expected non-empty string")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		t.Fatalf("invalid uuid string: %v", err)
	}
}
