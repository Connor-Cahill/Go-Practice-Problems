package hamming

import "errors"

//Distance calculates the hamming of 2 DNA Strands
func Distance(a, b string) (int, error) {
	counter := 0
	if len(a) != len(b) {
		err := errors.New("DNA Strands must be the same length")
		return 0, err
	}
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
