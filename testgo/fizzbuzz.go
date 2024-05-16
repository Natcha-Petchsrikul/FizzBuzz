package main

import "strconv"

type FizzBuzz struct {
	input  int
	result string
}

func NewFizzBuzz(input int) FizzBuzz {
	fizzbuzz := FizzBuzz{}
	fizzbuzz.input = input
	fizzbuzz.result = ""
	return fizzbuzz
}
func (f *FizzBuzz) Check(n int) string {
	f.result = strconv.Itoa(n)
	f.result = CheckFizz(f.result)
	return f.result
}

func CheckFizz(s string) string {
	n, _ := strconv.Atoi(s)
	result := (map[bool]string{true: "Fizz", false: s})[n%3 == 0]
	return result
}
