package pangram

import "unicode"

//IsPangram takes in a string and returns bool if pangram
func IsPangram(sentence string) bool {
	alphaSet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	for _, c := range sentence {
		for i, x := range alphaSet {
			if unicode.ToLower(c) == x {
				alphaSet[i] = alphaSet[len(alphaSet)-1]
				alphaSet = alphaSet[:len(alphaSet)-1]
			}
		}
	}

	if len(alphaSet) == 0 {
		return true
	}
	return false
}
