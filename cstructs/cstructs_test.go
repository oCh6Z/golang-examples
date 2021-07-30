package cstructs

import "testing"

func TestHelloworld(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
