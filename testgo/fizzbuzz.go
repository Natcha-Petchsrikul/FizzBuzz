package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = checkFizz(result)
	result = checkBuzz(result)
	return result
}

func checkFizz(s string) string {
	n, _ := strconv.Atoi(s)
	return (map[bool]string{true: "Fizz", false: s})[n != 0 && n%3 == 0]
}

func checkBuzz(s string) string {
	n, _ := strconv.Atoi(s)
	return (map[bool]string{true: "Buzz", false: s})[n != 0 && n%5 == 0]
}
