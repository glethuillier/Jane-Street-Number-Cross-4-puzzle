package constraints

import (
	"testing"
)

func TestPredicates(t *testing.T) {
	testCases := []struct {
		rowID    uint
		numbers  []uint
		expected bool
	}{
		// squares (row 0)
		{
			rowID:    0,
			numbers:  []uint{4, 9},
			expected: true,
		},
		{
			rowID:    0,
			numbers:  []uint{3, 9},
			expected: false,
		},

		// 1 more than a palindrome (row 1)
		{
			rowID:    1,
			numbers:  []uint{12, 1111111111111112},
			expected: true,
		},
		{
			rowID:    1,
			numbers:  []uint{11},
			expected: false,
		},

		// primes raised to primes (row 2)
		{
			rowID:    2,
			numbers:  []uint{8, 823543},
			expected: true,
		},
		{
			rowID:    2,
			numbers:  []uint{8, 10},
			expected: false,
		},

		// sum digits equals 7 (row 3)
		{
			rowID:    3,
			numbers:  []uint{1111111, 1231},
			expected: true,
		},
		{
			rowID:    3,
			numbers:  []uint{111111, 1231},
			expected: false,
		},

		// fibonacci numbers (row 4)
		{
			rowID:    4,
			numbers:  []uint{377, 610},
			expected: true,
		},
		{
			rowID:    4,
			numbers:  []uint{377, 609},
			expected: false,
		},

		// squares (row 5)
		{
			rowID:    5,
			numbers:  []uint{4, 9},
			expected: true,
		},
		{
			rowID:    5,
			numbers:  []uint{3, 9},
			expected: false,
		},

		// %37 (row 6)
		{
			rowID:    6,
			numbers:  []uint{3626, 36963, 36999963},
			expected: true,
		},
		{
			rowID:    6,
			numbers:  []uint{3627},
			expected: false,
		},

		// palindrome %23 (row 7)
		{
			rowID: 7,
			numbers: []uint{
				161,
				414,
				575,
				828,
				989,
				1771,
				4554,
				7337,
			},
			expected: true,
		},
		{
			rowID: 7,
			numbers: []uint{
				161,
				414,
				500,
				828,
				989,
				1771,
				4554,
				7337,
			},
			expected: false,
		},

		// product ends with 1 (row 8)
		{
			rowID: 8,
			numbers: []uint{
				1779,
				1797,
				1919,
				1933,
				1977,
				1991,
			},
			expected: true,
		},
		{
			rowID: 8,
			numbers: []uint{
				1779,
				1797,
				1919,
				1933,
				1977,
				1990,
			}, expected: false,
		},

		// %88 (row 9)
		{
			rowID:    9,
			numbers:  []uint{87912, 87999912, 879999999912},
			expected: true,
		},
		{
			rowID:    9,
			numbers:  []uint{87911, 87999912, 879999999912},
			expected: false,
		},

		// 1 less than a palindrome (row 10)
		{
			rowID:    10,
			numbers:  []uint{10, 1111111111111110},
			expected: true,
		},
		{
			rowID:    10,
			numbers:  []uint{11},
			expected: false,
		},
	}

	for _, tc := range testCases {
		got := predicates[tc.rowID](tc.numbers)
		if got != tc.expected {
			t.Errorf("Expected %v for numbers %v but got %v", tc.expected, tc.numbers, got)
		}
	}
}
