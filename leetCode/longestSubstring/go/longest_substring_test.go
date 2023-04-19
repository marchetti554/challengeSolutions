package _go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	result := lengthOfLongestSubstring("dvdf")
	fmt.Println("Longest substring: ", result)
}
