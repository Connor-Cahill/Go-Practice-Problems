package strand

import (
	"fmt"
	"strconv"
)

//ToRNA takes a strand of DNA as a string and returns the RNA match
func ToRNA(dna string) string {
	var rnaStrand string
	for _, l := range dna {
		switch strconv.QuoteRune(l) {
		case "G":
			rnaStrand = rnaStrand + "C"
		case "C":
			rnaStrand = rnaStrand + "G"
		case "T":
			rnaStrand = rnaStrand + "A"
		case "A":
			rnaStrand = rnaStrand + "U"
		}
	}
	fmt.Printf("rnaStrand: %s", rnaStrand)
	return rnaStrand
}
