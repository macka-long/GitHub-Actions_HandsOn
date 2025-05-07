package main

import "testing"

func TestGetHelloMessage(t *testing.T) {
	want := "Hello, Go Action"
	got := GetHelloMessage()

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
