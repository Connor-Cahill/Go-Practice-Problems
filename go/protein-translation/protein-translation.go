package protein

import (
	"fmt"
	"strings"
)

//FromCodon takes in an RNA Codon as a string and returns the protein
func FromCodon(codon string) string {
	var protein string
	s := strings.ToLower(codon)
	switch {
	case s == "aug":
		protein = "Methionine"
	case s == "uuu" || s == "uuc":
		protein = "Phenylalanine"
	case s == "uua" || s == "uug":
		protein = "Leucine"
	case s == "ucu" || s == "ucc" || s == "uca" || s == "ucg":
		protein = "Serine"
	case s == "uau" || s == "uac":
		protein = "Tyrosine"
	case s == "ugu" || s == "ugc":
		protein = "Cysteine"
	case s == "ugg":
		protein = "Tryptophan"
	case s == "uaa" || s == "uag" || s == "uga":
		protein = ""
	}
	return protein
}

//FromRNA takes in an RNA strand and returns slice of proteins
func FromRNA(strand string) []string {
	var proteinSlice []string

	for i := 1; i <= len(strand); i = i + 3 {
		codon := fmt.Sprintf("%s%s%s", strand[i], strand[i+1], strand[i+2])
		protein := FromCodon(codon)
		proteinSlice = append(proteinSlice, protein)
	}
	return proteinSlice
}
