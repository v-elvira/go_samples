package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup tests...")
	start := time.Now()
	code := m.Run() // <----- run tests here
	fmt.Println("Teardown tests...")
	end := time.Now()
	fmt.Println("Tests took", end.Sub(start))
	os.Exit(code)
}

func TestIntMin(t *testing.T) {
	got := IntMin(2, -2)
	want := -2
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
