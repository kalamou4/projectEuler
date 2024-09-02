package euler

import "math"


//return all primes factor for a number 
func PrimeFactors(n int) map[int]int {
	factors := make(map[int]int)
	primes := PrimeSieve(n)

	for _, prime := range primes {
		for n%prime == 0 {
			factors[prime]++
			n /= prime
		}

	}
	return factors
}

// max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// smallestMultiple finds the smallest number divisible by all numbers from 1 to n
func SmallestMultiple(n int) int {
	allFactors := make(map[int]int)

	for i := 2; i <= n; i++ {
		factors := PrimeFactors(i)
		for prime, count := range factors {
			allFactors[prime] = Max(allFactors[prime], count)
		}
	}

	result := 1
	for prime, count := range allFactors {
		result *= int(math.Pow(float64(prime), float64(count)))
	}

	return result
}
