package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = CheckFizz(result)
	return result
}

func CheckFizz(s string) string {
	var n, _ = strconv.Atoi(s)
	result := (map[bool]string{true: "Fizz", false: s})[n%3 == 0]
	return result
}
