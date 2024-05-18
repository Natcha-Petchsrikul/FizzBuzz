package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = check(result, "Fizz", n%3 == 0)
	result = check(result, "Buzz", n%5 == 0)
	result = check(result, "FizzBuzz", n%15 == 0)
	return result
}

func check(original string, s string, c bool) string {
	if c {
		return s
	} else {
		return original
	}
}
