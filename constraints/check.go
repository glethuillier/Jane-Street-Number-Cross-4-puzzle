package constraints

import (
	"fmt"
	"math"
)

// highest possible number per row:
// 99,999,999,999 (11 digits)
const maxDigitsBoundNumberPerRow = 1e12 - 1

var (
	squaresNumbers              map[uint]struct{}
	primesRaisedToPrimesNumbers map[uint]struct{}
	fibonacciNumbers            map[uint]struct{}
)

func sieveOfEratosthenes(n int) (primes []int) {
	b := make([]bool, n)
	for i := 2; i < n; i++ {
		if b[i] {
			continue
		}

		primes = append(primes, i)

		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}

	return primes
}

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	var (
		n1 = uint(0)
		n2 = uint(1)
	)

	for i := uint(1); i <= n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n1
}

func init() {
	// squares
	squaresNumbers = make(map[uint]struct{})

	for i := uint(1); i < 400_000; i++ {
		product := i * i

		if product > maxDigitsBoundNumberPerRow {
			break
		}

		squaresNumbers[product] = struct{}{}
	}

	fmt.Printf("Initialized %d square numbers\n", len(squaresNumbers))

	// primes raised to primes
	primesRaisedToPrimesNumbers = make(map[uint]struct{})

	for _, p1 := range sieveOfEratosthenes(320_000) {
		for _, p2 := range sieveOfEratosthenes(32) {
			pow := math.Pow(float64(p1), float64(p2))

			if pow > maxDigitsBoundNumberPerRow {
				continue
			}

			primesRaisedToPrimesNumbers[uint(pow)] = struct{}{}
		}
	}

	fmt.Printf("Initialized %d primes raised to primes\n", len(primesRaisedToPrimesNumbers))

	// Fibonacci numbers
	fibonacciNumbers = make(map[uint]struct{})

	for i := 0; i < 60; i++ {
		fib := fibonacci(uint(i))
		if fib > maxDigitsBoundNumberPerRow {
			fmt.Println(i)
			break
		}

		fibonacciNumbers[fib] = struct{}{}
	}

	fmt.Printf("Initialized %d Fibonacci numbers\n", len(fibonacciNumbers))
}

func isSquare(number uint) bool {
	_, ok := squaresNumbers[number]
	return ok
}

func isFibonacci(number uint) bool {
	_, ok := fibonacciNumbers[number]
	return ok
}

func isPalindrome(n uint) bool {
	original := n
	reversed := uint(0)
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	return original == reversed
}

func isPrimeRaisedToPrime(number uint) bool {
	_, ok := primesRaisedToPrimesNumbers[number]
	return ok
}

// sumOfDigitsEqualToTarget checks whether the sum of the digits of a given number
// is equal to the target value
func sumOfDigitsEqualToTarget(number, target uint) bool {
	sum := uint(0)

	for number > 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}

	return sum == target
}

// productDigitsEndsWithOne checks if the product of the digits of a given number
// ends with the digit 1
func productDigitsEndsWithOne(number uint) bool {
	product := uint(1)

	for number > 0 {
		digit := number % 10
		product *= digit
		number /= 10
	}

	return product%10 == 1
}
