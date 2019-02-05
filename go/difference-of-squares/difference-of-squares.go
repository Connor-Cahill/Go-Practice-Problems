package diffsquares

//SumOfSquare is the sum of squares in an array
func SumOfSquare(numbers []int) int {
	var ret int
	for _, n := range numbers {
		ret += n * n
	}
	return ret
}

//SquareOfSum sums up numbers in slice and then sqaures them
func SquareOfSum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total * total
}

//Difference gives the difference of squares
func Difference(numOne int, numTwo int) int {
	return numOne - numTwo
}
