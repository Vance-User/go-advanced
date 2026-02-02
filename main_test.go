package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
		errMsg  string
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "factorial of negative", input: -1, want: 0, wantErr: true, errMsg: "factorial is not defined for negative numbers"},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Factorial(%d) error=%v wantErr=%v", tt.input, err, tt.wantErr)
			}
			if err != nil && tt.errMsg != "" && err.Error() != tt.errMsg {
				t.Fatalf("Factorial(%d) errMsg=%q want %q", tt.input, err.Error(), tt.errMsg)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("Factorial(%d)=%d want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
		errMsg  string
	}{
		{name: "n less than 2", input: 1, want: false, wantErr: true, errMsg: "prime check requires number >= 2"},
		{name: "2 is prime", input: 2, want: true, wantErr: false},
		{name: "3 is prime", input: 3, want: true, wantErr: false},
		{name: "4 is not prime", input: 4, want: false, wantErr: false},
		{name: "17 is prime", input: 17, want: true, wantErr: false},
		{name: "25 is not prime", input: 25, want: false, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("IsPrime(%d) error=%v wantErr=%v", tt.input, err, tt.wantErr)
			}
			if err != nil && tt.errMsg != "" && err.Error() != tt.errMsg {
				t.Fatalf("IsPrime(%d) errMsg=%q want %q", tt.input, err.Error(), tt.errMsg)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("IsPrime(%d)=%v want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name    string
		base    int
		exp     int
		want    int
		wantErr bool
		errMsg  string
	}{
		{name: "base^0", base: 7, exp: 0, want: 1, wantErr: false},
		{name: "2^8", base: 2, exp: 8, want: 256, wantErr: false},
		{name: "5^3", base: 5, exp: 3, want: 125, wantErr: false},
		{name: "0^5", base: 0, exp: 5, want: 0, wantErr: false},
		{name: "negative exponent", base: 2, exp: -1, want: 0, wantErr: true, errMsg: "negative exponents not supported"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exp)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Power(%d,%d) error=%v wantErr=%v", tt.base, tt.exp, err, tt.wantErr)
			}
			if err != nil && tt.errMsg != "" && err.Error() != tt.errMsg {
				t.Fatalf("Power(%d,%d) errMsg=%q want %q", tt.base, tt.exp, err.Error(), tt.errMsg)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("Power(%d,%d)=%d want %d", tt.base, tt.exp, got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	c1 := MakeCounter(0)

	if got := c1(); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
	if got := c1(); got != 2 {
		t.Fatalf("expected 2, got %d", got)
	}

	c2 := MakeCounter(10)
	if got := c2(); got != 11 {
		t.Fatalf("expected 11, got %d", got)
	}

	// Ensure counters are independent
	if got := c1(); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "double", factor: 2, input: 5, want: 10},
		{name: "triple", factor: 3, input: 5, want: 15},
		{name: "zero factor", factor: 0, input: 100, want: 0},
		{name: "negative factor", factor: -2, input: 4, want: -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MakeMultiplier(tt.factor)
			if got := m(tt.input); got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, sub, get := MakeAccumulator(100)

	if got := get(); got != 100 {
		t.Fatalf("expected 100, got %d", got)
	}

	add(50)
	if got := get(); got != 150 {
		t.Fatalf("expected 150, got %d", got)
	}

	sub(30)
	if got := get(); got != 120 {
		t.Fatalf("expected 120, got %d", got)
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		op   func(int) int
		want []int
	}{
		{"square", []int{1, 2, 3, 4}, func(x int) int { return x * x }, []int{1, 4, 9, 16}},
		{"double", []int{-1, 0, 2}, func(x int) int { return x * 2 }, []int{-2, 0, 4}},
		{"negate", []int{5, -3}, func(x int) int { return -x }, []int{-5, 3}},
		{"empty", []int{}, func(x int) int { return x + 1 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]int(nil), tt.in...) // copy to ensure original isn't modified

			got := Apply(tt.in, tt.op)

			if len(got) != len(tt.want) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %v want %v", got, tt.want)
				}
			}

			// Ensure input slice not modified
			for i := range tt.in {
				if tt.in[i] != original[i] {
					t.Fatalf("input slice modified: got %v want %v", tt.in, original)
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		pred func(int) bool
		want []int
	}{
		{"evens", []int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x%2 == 0 }, []int{2, 4, 6}},
		{"positives", []int{-3, -1, 0, 2, 5}, func(x int) bool { return x > 0 }, []int{2, 5}},
		{"gt10", []int{9, 10, 11, 20}, func(x int) bool { return x > 10 }, []int{11, 20}},
		{"none", []int{1, 3, 5}, func(x int) bool { return x%2 == 0 }, []int{}},
		{"empty", []int{}, func(x int) bool { return x > 0 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.in, tt.pred)

			if len(got) != len(tt.want) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %v want %v", got, tt.want)
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		init int
		op   func(int, int) int
		want int
	}{
		{"sum", []int{1, 2, 3, 4}, 0, func(a, c int) int { return a + c }, 10},
		{"product", []int{1, 2, 3, 4}, 1, func(a, c int) int { return a * c }, 24},
		{"max", []int{5, 1, 9, 2}, -999, func(a, c int) int {
			if c > a {
				return c
			}
			return a
		}, 9},
		{"empty sum", []int{}, 0, func(a, c int) int { return a + c }, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.in, tt.init, tt.op)
			if got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	addTwo := func(x int) int { return x + 2 }
	double := func(x int) int { return x * 2 }

	tests := []struct {
		name string
		f    func(int) int
		g    func(int) int
		in   int
		want int
	}{
		{"addTwo(double(x))", addTwo, double, 5, 12},
		{"double(addTwo(x))", double, addTwo, 5, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Compose(tt.f, tt.g)
			if got := h(tt.in); got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"swap positives", 5, 10, 10, 5},
		{"swap with zero", 0, 7, 7, 0},
		{"swap negatives", -1, -9, -9, -1},
		{"swap same", 3, 3, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Fatalf("SwapValues(%d,%d)=(%d,%d) want (%d,%d)",
					tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"swap positives", 5, 10, 10, 5},
		{"swap with zero", 0, 7, 7, 0},
		{"swap negatives", -1, -9, -9, -1},
		{"swap same", 3, 3, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.a
			b := tt.b
			SwapPointers(&a, &b)
			if a != tt.wantA || b != tt.wantB {
				t.Fatalf("after SwapPointers a=%d b=%d want a=%d b=%d",
					a, b, tt.wantA, tt.wantB)
			}
		})
	}
}
