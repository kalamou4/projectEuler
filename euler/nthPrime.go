package euler

import (
	"math"
)

// takes an int and returns
func PrintNthPrime(n int) int {

	var value, count = 0, 0
	for count < n {
		value++
		if IsPrime(value) {
			count++
		}
	}

	return value
}

func IsPrime(n int) bool {
	if n >= 2 {

		if n == 2 || n == 3 {
			return true
		}

		square_root := int(math.Sqrt(float64(n)))
		for i := 2; i <= square_root; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
