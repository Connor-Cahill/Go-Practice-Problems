package reverse

//String takes in a string and reverses it
func String(word string) string {
	if word == "" {
		return ""
	}
	chars := []rune(word)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
