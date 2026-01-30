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
