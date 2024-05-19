package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = check(n, result)
	return result
}

func check(n int, original string) string {
	return (map[bool]string{true: "Fizz", false: original})[n%3 == 0]
}
