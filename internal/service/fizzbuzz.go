package service

import "strconv"

func FizzBuzz(num int) string {
	switch {
	case num < 0:
		return "The number must be positive."
	case num%3 == 0 && num%5 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(num)
	}
}
