package euler


// The function LargestPrimeFactor returns the largest prime factor of a given limit.
func LargestPrimeFactor(limit int) int {
	// Start with the smallest prime factor
	prime := 2

	// This part of the code is a loop that iterates as long as the square of the current prime number
	// (`prime*prime`) is less than or equal to the given limit.
	for prime*prime <= limit {
		if limit%prime == 0 {
			limit /= prime
		} else {
			prime++
		}
	}
	return limit
}
