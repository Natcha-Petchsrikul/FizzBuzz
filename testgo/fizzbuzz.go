package main

import (
	"strconv"
)

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = CheckFizzBuzz(result)
	result = CheckFizz(result)
	result = CheckBuzz(result)
	return result
}

func CheckFizz(s string) string {
	var n, _ = strconv.Atoi(s)
	result := (map[bool]string{true: "Fizz", false: s})[n != 0 && n%3 == 0]
	return result
}

func CheckBuzz(s string) string {
	var n, _ = strconv.Atoi(s)
	result := (map[bool]string{true: "Buzz", false: s})[n != 0 && n%5 == 0]
	return result
}

func CheckFizzBuzz(s string) string {
	var n, _ = strconv.Atoi(s)
	result := (map[bool]string{true: "FizzBuzz", false: s})[n != 0 && n%15 == 0]
	return result
}
