package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	msg := GetMessage()
	want := "Hello 🗺️ !"
	if msg != want {
		t.Fatalf(`GetMessage() = %q, want match for %#q`, msg, want)
	}
}
