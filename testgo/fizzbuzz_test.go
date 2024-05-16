package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Check(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Check(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
