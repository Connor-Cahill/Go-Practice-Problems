package isogram

import "unicode"

//IsIsogram takes in a string and returns true or false on whether is is a isogram
func IsIsogram(word string) bool {
	if word == "" {
		return true
	}
	letters := make(map[rune]int)
	for _, c := range word {
		_, ok := letters[unicode.ToUpper(c)]
		if ok {
			return false
		}
		if unicode.IsLetter(c) {
			letters[unicode.ToUpper(c)] = 1
		}
	}
	return true
}
