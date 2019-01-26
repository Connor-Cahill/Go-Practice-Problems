package raindrops

import "strconv"

//Convert changes int to certain string based on num
func Convert(num int) string {
	var s string
	factors := [3]int{3, 5, 7}
	rainNoises := [3]string{"Pling", "Plang", "Plong"}

	for i := range factors {
		if num%factors[i] == 0 {
			s = s + rainNoises[i]
		}
	}
	if s == "" {
		newS := strconv.Itoa(num)
		return newS
	}

	return s
}
