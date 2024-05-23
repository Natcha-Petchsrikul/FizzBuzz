package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = checkFizzBuzz(result)
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

func checkFizzBuzz(s string) string {
	n, _ := strconv.Atoi(s)
	return (map[bool]string{true: "FizzBuzz", false: s})[n != 0 && n%15 == 0]
}
