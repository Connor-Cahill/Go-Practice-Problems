package etl

import "strings"

//Transform takes a map and updates its format
func Transform(oldMap map[int][]string) map[string]int {
	newMap := make(map[string]int)

	for k, v := range oldMap {
		for _, c := range v {
			newMap[strings.ToLower(c)] = k
		}
	}
	return newMap
}
