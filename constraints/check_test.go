package constraints

import (
	"math"
	"testing"
)

func TestIsSquare(t *testing.T) {
	testCases := []struct {
		input    uint
		expected bool
	}{
		{4, true},
		{5, false},
		{9, true},
		{10, false},
		{16, true},
		{25, true},
	}

	for _, tc := range testCases {
		got := isSquare(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v for input %d but got %v", tc.expected, tc.input, got)
		}
	}
}

func TestIsFibonacci(t *testing.T) {
	testCases := []struct {
		input    uint
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, false},
		{8, true},
		{9, false},
		{1000000, false},
	}

	for _, tc := range testCases {
		got := isFibonacci(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v for input %d but got %v", tc.expected, tc.input, got)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    uint
		expected bool
	}{
		{0, true},
		{1, true},
		{121, true},
		{123, false},
		{12121, true},
		{123456789654321, false},
		{13131, true},
		{156156, false},
	}

	for _, tc := range testCases {
		got := isPalindrome(tc.input)
		if got != tc.expected {
			t.Errorf("Expected %v for input %d but got %v", tc.expected, tc.input, got)
		}
	}
}

func TestIsPrimeRaisedToPrime(t *testing.T) {
	testCases := []struct {
		base     float64
		exponent float64
		expected bool
	}{
		{2, 2, true},
		{2, 3, true},
		{2, 4, false},
		{3, 2, true},
		{3, 3, true},
		{5, 2, true},
		{7, 2, true},
		{11, 2, true},
		{13, 2, true},
		{17, 2, true},
		{19, 2, true},
		{19, 8, false},
		{2, 7, true},
		{3, 5, true},
	}

	for _, tc := range testCases {
		expected := isPrimeRaisedToPrime(uint(math.Pow(tc.base, tc.exponent)))
		if expected != tc.expected {
			t.Errorf("Expected %f^%f to be a prime composition but it is not", tc.base, tc.exponent)
		}
	}
}
