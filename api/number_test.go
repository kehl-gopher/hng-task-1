package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test if a number is armstrong or narcissitic
func TestArmStrong(t *testing.T) {

	n := Parity(370)
	assert.True(t, n.checkNumberIsArmStrong(), "You narcissitic bitch")

	b := Parity(300)
	assert.False(t, b.checkNumberIsArmStrong(), "You are not narcissitic I guess")
}

// check for the parity of a  number
func TestIfNumberIsEvenOrOdd(t *testing.T) {
	e := Parity(100)
	assert.Equal(t, "even", e.checkParity(), "Even number you just bloody evil")
	o := Parity(99)

	assert.Equal(t, "odd", o.checkParity(), "Odd number you ugly asf")
}

func TestNumberIsPrime(t *testing.T) {

	isPrime := Parity(10)
	assert.False(t, isPrime.checkIsPrime(), "why the fuck are you prime")

	t.Run("Check if 0 return not prime", func(t *testing.T) {
		isPrime = Parity(0)
		assert.False(t, isPrime.checkIsPrime(), nil)
	})

	t.Run("Check if 1 return not prime", func(t *testing.T) {
		isPrime = 1
		assert.False(t, isPrime.checkIsPrime(), nil)
	})

	t.Run("Check if 3 is prime", func(t *testing.T) {
		isPrime = 3
		assert.True(t, isPrime.checkIsPrime(), "Why you not prime")
	})

	// check if 2 returned a prime
	t.Run("Check if 2 is prime", func(t *testing.T) {
		isPrime = 2
		assert.True(t, isPrime.checkIsPrime(), "Dude you meant to be prime")
	})

	t.Run("check 7 is prime", func(t *testing.T) {

		isPrime = 7
		assert.True(t, isPrime.checkIsPrime(), nil)
	})

	t.Run("check 100 is not prime", func(t *testing.T) {

		isPrime = 100
		assert.False(t, isPrime.checkIsPrime(), nil)
	})

	t.Run("check 47 is prime", func(t *testing.T) {
		isPrime = 47
		assert.True(t, isPrime.checkIsPrime(), nil)
	})
}
