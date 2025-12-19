package greet

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tester")
	want := "Hello, Tester!"
	if got != want {
		t.Fatalf("Hello() = %q; want %q", got, want)
	}
}
