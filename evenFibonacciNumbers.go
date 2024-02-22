package projecteuler

func EvenFibonacciNumbers() []int {
	var i = 3
	var tab = []int{}

	for Fibonacci(i-1) + Fibonacci(i-2) < 4000000 {
		sum := Fibonacci(i-1) + Fibonacci(i-2)
		if sum % 2 == 0 {
			tab = append(tab, sum)
		}
	}
	return tab	
}

func Fibonacci(x int) int{

	for x >= 3 {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
	return 0
}