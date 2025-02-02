package main

import (
	"math"
)

// check if a number is even or odd
func (a Parity) checkParity() string {

	if a%2 == 0 {
		return "even"
	}
	return "odd"
}

// check if a number is armstrong
func (a Parity) checkNumberIsArmStrong() bool {
	numOfDigits := Parity(math.Log10(float64(a))) + 1 // give me the length of the digits inputed
	sum := Parity(0)
	temp := a

	for temp > 0 {
		digit := temp % 10
		sum += Parity(math.Pow(float64(digit), float64(numOfDigits)))
		temp /= 10
	}

	return sum == a
}

// using the trial division algorihtm cause that bitch the easiest to use Don't Mind me
func (a Parity) checkIsPrime() bool {
	if a <= 1 {
		return false
	}

	if a <= 3 {
		return true
	}
	if a%2 == 0 || a%3 == 0 {
		return false
	}

	// check only odd number and skip all multiples of 3
	for i := Parity(5); i*i <= a; i += 6 {

		if a%i == 0 || a%(i+2) == 0 {
			return false
		}
	}
	return true
}
