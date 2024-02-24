package projecteuler


func MultipleOfFiveOrThree(three, five, x int) int {
	var sum int
	var i = 1
	for i < x {
		if i%three == 0 || i%five == 0 {
			sum += i
		}
		i++
	}
	return sum
}
