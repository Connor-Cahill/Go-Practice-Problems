package wordcount

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

//Frequency is a map of words with their frequency in a phrase (histogram)
type Frequency map[string]int

//WordCount takes in a phrase and returns map with freq of each word in phrase
func WordCount(phrase string) Frequency {
	// ret := make(Frequency)
	ret := make(Frequency)
	phrase = strings.Replace(phrase, "\n", "", len(phrase))
	phrase = strings.Replace(phrase, ",", " ", len(phrase))
	phrase = strings.TrimSpace(phrase)
	reg, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		panic(err)
	}
	phrase = reg.ReplaceAllString(phrase, " ")
	words := strings.Split(phrase, " ")
	fmt.Println(words)
	for _, w := range words {
		if w == "" {
			break
		}
		if unicode.IsLetter(rune(w[0])) != true && unicode.IsNumber(rune(w[0])) != true {
			w = strings.Replace(w, "'", "", len(w))
		}
		if _, found := ret[strings.ToLower(w)]; found {
			ret[strings.ToLower(w)]++
		} else {
			ret[strings.ToLower(w)] = 1
		}
	}
	return ret
}
