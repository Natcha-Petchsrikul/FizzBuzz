package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = check(result, "Fizz", n%3 == 0)
	result = check(result, "Buzz", n%5 == 0)
	return result
}

func check(original, s string, c bool) string {
	if c {
		return s
	}
	return original
}
