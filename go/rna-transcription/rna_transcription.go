package strand

//ToRNA takes a strand of DNA as a string and returns the RNA match
func ToRNA(dna string) string {
	var rnaStrand string
	for _, l := range dna {
		switch l {
		case 'G':
			rnaStrand += "C"
		case 'C':
			rnaStrand += "G"
		case 'T':
			rnaStrand += "A"
		case 'A':
			rnaStrand += "U"
		}
	}
	return rnaStrand
}
