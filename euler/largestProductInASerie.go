package euler

func LargestProductInASerie(numbers []int) int {

	maxProduct := 0
	for i := range numbers {
		if i+5 <= len(numbers) {
			product := ArrayProduct(numbers[i:i+5])
			if product > maxProduct{
				maxProduct = product
			}
		}
	}
	return maxProduct
}

func ArrayProduct(arr []int) int {
	product := 1
	for _,val := range arr{
		product *= val
	}
	return product
}
