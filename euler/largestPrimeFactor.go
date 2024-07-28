package euler

import (
	"math"
)

// Generate sends the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Filter copies the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}

// Sieve returns a slice of primes up to the given limit.
func Sieve(limit int) []int {
	//A prime factor or "limit" must be lower than its square root
	limit = int(math.Sqrt(float64(limit)))
	ch := make(chan int)
	go Generate(ch)

	primes := []int{}
	for {
		prime := <-ch
		if prime > limit {
			break
		}
		primes = append(primes, prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}

// LargestPrime returns an int reprensenting a the highest prime factor
func LargestPrime(limit int) int {
	primes := Sieve(limit)
	for i := len(primes) - 1; i >= 0; i-- {
		if limit%primes[i] == 0 {
			return primes[i]
		}
	}
	return limit
}
