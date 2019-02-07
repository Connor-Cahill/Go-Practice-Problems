package luhn

import (
	"fmt"
	"unicode"
)

//Valid takes in a CC number and returns bool if valid
func Valid(number string) bool {
	var sum int
	for _, n := range number {
		if unicode.IsSpace(n) {
			fmt.Println("space")
			break
		}
		if unicode.IsLetter(n) {
			fmt.Println("Letter")
			return false
		}
		fmt.Println(int(n))
		tempNum := int(n) * 2
		if int(n) > 9 {
			tempNum -= 9
		}
		sum += tempNum
	}
	fmt.Println(sum)
	if sum%10 == 0 {
		return true
	}
	return false
}
