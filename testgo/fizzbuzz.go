package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = checkFizz(result)
	return result
}

func checkFizz(s string) string {
	n, _ := strconv.Atoi(s)
	return (map[bool]string{true: "Fizz", false: s})[n%3 == 0]
}
