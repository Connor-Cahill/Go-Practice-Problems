package accumulate

import "strings"

type convert func(string) string

//Accumulate takes in slice of words and command and performs command
func Accumulate(words []string, command convert) []string {
	var newWords []string
	for _, w := range words {
		strings.ToLower(w)
		newWords = append(newWords, command(w))
	}
	return newWords
}
