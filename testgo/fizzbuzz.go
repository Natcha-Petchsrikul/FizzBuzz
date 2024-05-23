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
	m := map[int]string{3: "Fizz", 5: "Buzz"}
	a := [2]int{3, 5}
	for _, v := range a {
		if f.Input%v == 0 {
			result = m[v]
		}
	}
	return result
}
