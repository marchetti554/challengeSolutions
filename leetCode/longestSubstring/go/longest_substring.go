package _go

import "strings"

// Incomplete
func lengthOfLongestSubstring(s string) int {
	var possibleLongest, longest string
	chars := []rune(s)
	if len(chars) == 0 {
		return 0
	}
	if len(chars) == 1 {
		return 1
	}
	for i := 0; i < len(chars); i++ {
		as := string(chars[i])
		if strings.Contains(possibleLongest, as) {
			possibleLongest = as
		} else {
			possibleLongest = possibleLongest + as
		}
		if len(possibleLongest) > len(longest) {
			longest = possibleLongest
		}
	}
	return len(longest)
}
