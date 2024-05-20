package main

import "strconv"

type FizzBuzz struct {
	input  int
	Result string
}

func NewFizzBuzz(input int) FizzBuzz {
	fizzbuzz := FizzBuzz{}
	fizzbuzz.input = input
	fizzbuzz.calculate()
	return fizzbuzz
}
func (f *FizzBuzz) calculate() {
	f.Result = strconv.Itoa(f.input)
	m := map[int]string{3: "Fizz", 5: "Buzz", 15: "FizzBuzz"}
	a := [3]int{3, 5, 15}
	for _, v := range a {
		if f.input%v == 0 {
			f.Result = m[v]
		}
	}
}
