package protein

import (
	"fmt"
	"strings"
)

//FromCodon takes in an RNA Codon as a string and returns the protein
func FromCodon(codon string) (string, error) {
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
	return protein, nil
}

//FromRNA takes in an RNA strand and returns slice of proteins
func FromRNA(strand string) ([]string, error) {
	var proteinSlice []string

	for i := 0; i <= len(strand)-3; i = i + 3 {
		codon := string(strand[i]) + string(strand[i+1]) + string(strand[i+2])
		fmt.Println(codon)
		protein, err := FromCodon(string(codon))
		if err != nil {
			panic(err)
		}
		if protein == "" {
			return proteinSlice, nil
		}
		proteinSlice = append(proteinSlice, protein)
	}
	return proteinSlice, nil
}
