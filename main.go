package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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

// Part 3: Higher-Order Functions

// Apply returns a new slice with operation applied to each element.
// It does not modify the original slice.
func Apply(nums []int, operation func(int) int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = operation(n)
	}
	return out
}

// Filter returns a new slice containing only elements where predicate returns true.
func Filter(nums []int, predicate func(int) bool) []int {
	out := make([]int, 0)
	for _, n := range nums {
		if predicate(n) {
			out = append(out, n)
		}
	}
	return out
}

// Reduce reduces a slice to a single value using the operation function.
func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	acc := initial
	for _, n := range nums {
		acc = operation(acc, n)
	}
	return acc
}

// Compose returns a new function h(x) = f(g(x)).
func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// Part 4: Process Explorer
//
// A process is a running program. Each process has a unique ID called a PID.
// Process isolation is important because it prevents other processes from directly
// reading or modifying this program's memory (security + stability).
//
// NOTE: A slice in Go has a small "header" (pointer/len/cap). &data is the address
// of that header variable, while &data[0] is the address of the first element in
// the underlying array.

func ExploreProcess() {
	fmt.Println("====== Process Information ======")

	pid := os.Getpid()
	ppid := os.Getppid()

	fmt.Printf("Current Process ID: %d\n", pid)
	fmt.Printf("Parent Process ID: %d\n", ppid)

	data := []int{1, 2, 3, 4, 5}

	fmt.Printf("Memory address of slice (header variable): %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	fmt.Println("Note: Other processes cannot access these addresses due to process isolation.")
	fmt.Println()
}
