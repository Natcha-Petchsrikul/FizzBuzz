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
func (f FizzBuzz) Check(n int) string {
	return strconv.Itoa(f.input)
}
