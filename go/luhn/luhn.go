package luhn

import (
	"fmt"
	"math"
)

//Valid takes in a CC number and returns bool if valid
func Valid(number string) bool {
	//	sum of all the numbers after modified
	var sum int
	if len(number) < 2 {
		return false
	}
	sum = doubleNum(number)
	if sum%10 == 0 {
		return true
	}
	return false
}

func doubleNum(number string) int {
	var sum int
	var i float64
	length := float64(len(number))
	for i = 0; math.Abs(i) < length; i -= 2 {
		tempNum := (int(number[int(i)]) - '0') * 2
		fmt.Println(tempNum)
		if tempNum > 9 {
			tempNum -= 9
		}
		sum += tempNum
		sum += (int(number[int(i+1)]) - '0')
	}
	return sum
}
