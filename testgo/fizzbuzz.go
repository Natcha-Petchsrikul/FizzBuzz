package main

import "strconv"

func FizzBuzz(n int) string {
	result := strconv.Itoa(n)
	result = check(result, n)
	return result
}

func check(original string, n int) string {
	m := map[int]string{3: "Fizz", 5: "Buzz", 15: "FizzBuzz"}
	a := [3]int{15, 3, 5}
	for _, v := range a {
		if n%v == 0 {
			return m[v]
		}
	}
	return original
}
