package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = (map[bool]string{true: "Fizz", false: result})[n%3 == 0]
	result = (map[bool]string{true: "Buzz", false: result})[n%5 == 0]

	return result
}
