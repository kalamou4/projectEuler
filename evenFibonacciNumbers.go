package projecteuler

func Fibonacci(index int) int {

	if index == 1 || index == 2 {
		return 1
	} else if index > 2  {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
	return 0
}


func EvenTermFibonnaci(limit int) int{

	var total, i int
	for Fibonacci(i) < limit {
		if Fibonacci(i) % 2 == 0 {
			total += Fibonacci(i)
		} 
		i++
	}
	return total
}
