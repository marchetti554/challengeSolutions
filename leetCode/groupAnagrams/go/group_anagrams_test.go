package _go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(input)
	fmt.Println(result)
}
