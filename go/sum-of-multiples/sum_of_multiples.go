package summultiples

//SumMultiples takes in divisors and a limit and returns sum off all multiples of divisor in limit
func SumMultiples(limit int, divisors []int) int {
	var sum int                  //	Sum that will be returned
	for i := 0; i < limit; i++ { // iterate over numbers in range of limit
		for _, d := range divisors {
			if d == 0 { //	if divisor = 0 break the loop (can't divide by 0)
				break
			}
			if i%d == 0 { //	if i divides evenly by divisor then add it to the sum
				sum += i
				break
			}
		}
	}
	return sum
}
