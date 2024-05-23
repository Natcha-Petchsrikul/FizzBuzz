package main

import "strconv"

type FizzBuzz struct {
	Input  int
	Result string
}

func NewFizzBuzz(input int) FizzBuzz {
	fizzbuzz := FizzBuzz{}
	fizzbuzz.Input = input
	fizzbuzz.Result = fizzbuzz.calculate()
	return fizzbuzz
}

func (f *FizzBuzz) calculate() string {
	result := strconv.Itoa(f.Input)
	result = map[bool]string{true: "Fizz", false: result}[f.Input%3 == 0]
	return result
}
