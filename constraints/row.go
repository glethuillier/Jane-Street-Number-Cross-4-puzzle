package constraints

// Rows-Related Constraints

// predicates correspond to the rows' constraints
// “Each row has been supplied with a clue.“
// “Every number formed by concatenating consecutive groups of unshaded cells within a row
// must satisfy the clue given for the row.“
var predicates = map[uint]func([]uint) bool{
	0: predicateSquare(),
	// “1 more than a palindrome“:
	// decrement the number and check whether it is a palindrome:
	1: predicatePalindrome(-1),
	2: predicatePrimeRaisedToPrime(),
	3: predicateSumsEqualsSeven(),
	4: predicateFibonacci(),
	5: predicateSquare(),
	6: predicateMultiple(37),
	7: predicatePalindromeMod23(),
	8: predicateProductDigitEndsWithOne(),
	9: predicateMultiple(88),
	// “1 less than a palindrome“:
	// increment the number and check whether it is a palindrome:
	10: predicatePalindrome(+1),
}

func NumbersSatisfyTheRowClue(rowID uint, numbers []uint) bool {
	return predicates[rowID](numbers)
}

func predicatePalindrome(delta int) func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !isPalindrome(uint(int(n) + delta)) {
				return false
			}
		}

		return true
	}
}

func predicatePrimeRaisedToPrime() func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !isPrimeRaisedToPrime(n) {
				return false
			}
		}

		return true
	}
}

func predicateSumsEqualsSeven() func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !sumOfDigitsEqualToTarget(n, 7) {
				return false
			}
		}

		return true
	}
}

func predicateFibonacci() func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !isFibonacci(n) {
				return false
			}
		}

		return true
	}
}

func predicateSquare() func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !isSquare(n) {
				return false
			}
		}

		return true
	}
}

func predicateMultiple(mod int) func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if n%uint(mod) != 0 {
				return false
			}
		}

		return true
	}
}

func predicatePalindromeMod23() func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !isPalindrome(uint(int(n))) {
				return false
			}
			if n%23 != 0 {
				return false
			}
		}

		return true
	}
}

func predicateProductDigitEndsWithOne() func([]uint) bool {
	return func(numbers []uint) bool {
		for _, n := range numbers {
			if !productDigitsEndsWithOne(n) {
				return false
			}
		}

		return true
	}
}
