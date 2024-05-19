package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = check("Fizz", result, n%3 == 0)
	result = check("Buzz", result, n%5 == 0)
	result = check("FizzBuzz", result, n%15 == 0)
	return result
}

func check(s string, original string, c bool) string {
	return (map[bool]string{true: s, false: original})[c]
}
