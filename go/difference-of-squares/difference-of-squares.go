package diffsquares

//SumOfSquares is the sum of squares of length numbers
func SumOfSquares(length int) int {
	var ret int
	for i := 0; i <= length; i++ {
		ret += i * i
	}
	return ret
}

//SquareOfSum sums up numbers in slice and then sqaures them
func SquareOfSum(length int) int {
	var total int
	for i := 0; i <= length; i++ {
		total += i
	}
	return total * total
}

//Difference gives the difference of squares
func Difference(length int) int {
	numOne := SumOfSquares(length)
	numTwo := SquareOfSum(length)

	return numTwo - numOne
}
