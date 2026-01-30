package main

import (
	"errors"
	"math"
)

// Part 1: Math Operations

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

// Part 2: Function Factory & Closures

// MakeCounter returns a function that increments and returns a counter.
// Each returned counter is independent because it captures its own state.
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// MakeMultiplier returns a function that multiplies input by the captured factor.
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// MakeAccumulator returns three functions that share the same captured state:
// add increases the accumulator, subtract decreases it, and get returns current value.
func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	value := initial

	add = func(x int) {
		value += x
	}
	subtract = func(x int) {
		value -= x
	}
	get = func() int {
		return value
	}

	return
}
