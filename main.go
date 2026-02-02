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

// Part 5: Pointer Playground & Escape Analysis

// DoubleValue takes an int by value.
// Will this modify the original variable? No.
// Why? Go passes arguments by value (it makes a copy).
func DoubleValue(x int) {
	x = x * 2
}

// DoublePointer takes a pointer to an int.
// Will this modify the original variable? Yes.
// Why? We modify the value at the memory address.
func DoublePointer(x *int) {
	*x = *x * 2
}

// CreateOnStack returns a value (not a pointer).
// This variable stays on the stack.
func CreateOnStack() int {
	n := 42
	return n
}

// CreateOnHeap returns a pointer to a local variable.
// This variable escapes to the heap.
func CreateOnHeap() *int {
	n := 99
	return &n
}

// SwapValues swaps two ints and returns the swapped values.
// This does not change the originals unless the caller assigns the returned values.
func SwapValues(a, b int) (int, int) {
	return b, a
}

// SwapPointers swaps the values that two pointers point to.
// This modifies the original variables.
func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

// AnalyzeEscape calls both CreateOnStack and CreateOnHeap.
// Run escape analysis with:
//
//	go build -gcflags '-m' main.go
//
// Add your explanation below after you run the command.
func AnalyzeEscape() {
	_ = CreateOnStack()
	_ = CreateOnHeap()
}

/*
Escape analysis explanation (from: go build -gcflags '-m' main.go):

  - CreateOnHeap(): the local variable `n` escapes to the heap because we return its address (&n).
    The output shows: "moved to heap: n". The value must stay valid after the function returns.

  - CreateOnStack(): returns an int value (not a pointer), so it typically does NOT need to escape.
    Returning a value copies it out, so the local variable can stay on the stack.

Other escapes in this project:
  - Closures: captured variables escape (ex: "moved to heap: count" and "moved to heap: value")
    because the returned function needs that state after the outer function returns.

"Escapes to heap" means the compiler decided a variable must be heap-allocated instead of stack-allocated,
usually because its address is returned/stored or it outlives the function call.
*/
func main() {
	// Part 4: Process Information
	ExploreProcess()

	// Part 1: Math Operations
	fmt.Println("====== Math Operations ======")

	for _, n := range []int{0, 5, 10} {
		f, err := Factorial(n)
		if err != nil {
			fmt.Printf("Factorial(%d) error: %v\n", n, err)
			continue
		}
		fmt.Printf("Factorial(%d) == %d\n", n, f)
	}

	for _, n := range []int{17, 20, 25} {
		isP, err := IsPrime(n)
		if err != nil {
			fmt.Printf("IsPrime(%d) error: %v\n", n, err)
			continue
		}
		fmt.Printf("IsPrime(%d) == %v\n", n, isP)
	}

	for _, p := range [][2]int{{2, 8}, {5, 3}} {
		result, err := Power(p[0], p[1])
		if err != nil {
			fmt.Printf("Power(%d,%d) error: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("Power(%d,%d) == %d\n", p[0], p[1], result)
	}

	fmt.Println()

	// Part 2: Closures
	fmt.Println("====== Closure Demonstration ======")

	counter1 := MakeCounter(0)
	counter2 := MakeCounter(100)

	fmt.Printf("Counter1: %d\n", counter1())
	fmt.Printf("Counter1: %d\n", counter1())
	fmt.Printf("Counter2: %d\n", counter2())
	fmt.Printf("Counter2: %d\n", counter2())
	fmt.Printf("Counter1: %d (independent)\n", counter1())

	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)
	fmt.Printf("Doubler(5) = %d\n", doubler(5))
	fmt.Printf("Tripler(5) = %d\n", tripler(5))

	add, sub, get := MakeAccumulator(100)
	add(50)
	sub(30)
	fmt.Printf("Accumulator result = %d\n", get())

	fmt.Println()

	// Part 3: Higher-Order Functions
	fmt.Println("====== Higher-Order Functions ======")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", nums)

	squared := Apply(nums, func(x int) int { return x * x })
	fmt.Printf("Squared:  %v\n", squared)

	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Evens:    %v\n", evens)

	sum := Reduce(nums, 0, func(acc, cur int) int { return acc + cur })
	fmt.Printf("Sum:      %d\n", sum)

	doubleThenAddTen := Compose(
		func(x int) int { return x + 10 }, // f
		func(x int) int { return x * 2 },  // g
	)
	fmt.Printf("Double then add 10 (5) = %d\n", doubleThenAddTen(5))

	fmt.Println()

	// Part 5: Pointers
	fmt.Println("====== Pointer Demonstration ======")

	a, b := 5, 10
	fmt.Printf("Before SwapValues: a=%d, b=%d\n", a, b)
	newA, newB := SwapValues(a, b)
	fmt.Printf("SwapValues returned: a=%d, b=%d (originals unchanged)\n", newA, newB)
	fmt.Printf("Original still: a=%d, b=%d\n", a, b)

	fmt.Printf("Before SwapPointers: a=%d, b=%d\n", a, b)
	SwapPointers(&a, &b)
	fmt.Printf("After SwapPointers: a=%d, b=%d (originals changed)\n", a, b)

	// Keep this call in the program (escape analysis demo)
	AnalyzeEscape()
}
