package euler

import "math"

//function returns the sum of the square of n numbers
func SumOfSquares(n int) int{
	var sum_of_squares_forumule = (n * ((2*n) + 1)*(n +1))/6
	return sum_of_squares_forumule

}

//function returs the square of the sum of n numbers
func SquareOfTheSum(n int) int{
	var square_of_the_sum = (n * (n + 1))/2
	return int(math.Pow(float64(square_of_the_sum), 2))
}

//return the defference between the square of the sum and sum of squares
func SquareDifference(n int) int {
	return SquareOfTheSum(n) - SumOfSquares(n)
}