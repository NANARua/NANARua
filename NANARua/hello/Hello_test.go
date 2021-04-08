package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello world"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

// run "go test in termial" to text the function
