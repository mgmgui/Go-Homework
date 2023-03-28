package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	want := 3
	if got := Add(1, 2); got != want {
		t.Errorf("Add() = %d, want %d", got, want)
	}
}
