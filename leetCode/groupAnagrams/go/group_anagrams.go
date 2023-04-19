package _go

import (
	"sort"
)

func groupAnagrams(strings []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, v := range strings {
		sortedString := sortString(v)
		if _, ok := anagramMap[sortedString]; ok {
			anagramMap[sortedString] = append(anagramMap[sortedString], v)
		} else {
			anagramMap[sortedString] = []string{v}
		}
	}

	v := make([][]string, 0, len(anagramMap))

	for _, value := range anagramMap {
		v = append(v, value)
	}

	return v
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
