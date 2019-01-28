package proverb

// Proverb takes in a slic of strings and returns a proverb
func Proverb(rhyme []string) []string {
	var proverb []string
	for i := range rhyme {
		if i != len(rhyme)-1 {
			proverb = append(proverb, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
		} else {
			proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
		}
	}
	return proverb
}
