package euler

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func PrimeSieve(limit int) []int {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.

	primes := []int{}
	for {
		prime := <-ch

		// Break the loop if prime exceeds the limit (for the usual case)
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