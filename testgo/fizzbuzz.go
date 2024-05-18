package main

import "strconv"

func FizzBuzz(n int) string {
	result := (map[bool]string{true: "Fizz", false: strconv.Itoa(n)})[n%3 == 0]
	return result
}
