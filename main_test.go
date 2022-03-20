package main

import (
	"testing"

	"github.com/skanehira/go-work/example/strings"
)

func TestToUpperError(t *testing.T) {
	want := "some error"
	_, got := strings.ToUpper("hello")

	if got == nil {
		t.Fatal("error is nil")
	}
	if want != got.Error() {
		t.Fatalf("unexpected error, want: %s, got: %s", want, got)
	}
}
