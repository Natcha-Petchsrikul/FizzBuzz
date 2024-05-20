package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput3(t *testing.T) {
	input := 3
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn4WhenInput4(t *testing.T) {
	input := 4
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "4"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnBuzzWhenInput5(t *testing.T) {
	input := 5
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "Buzz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput6(t *testing.T) {
	input := 6
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn7WhenInput7(t *testing.T) {
	input := 7
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "7"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn8WhenInput8(t *testing.T) {
	input := 8
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "8"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput9(t *testing.T) {
	input := 9
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnBuzzWhenInput10(t *testing.T) {
	input := 10
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "Buzz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn11WhenInput11(t *testing.T) {
	input := 11
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "11"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput12(t *testing.T) {
	input := 12
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn13WhenInput13(t *testing.T) {
	input := 13
	fizzbuzz := NewFizzBuzz(input)

	got := fizzbuzz.Result

	want := "13"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
