package _go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(checkMagazine([]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "today"}))          //Yes
	fmt.Println(checkMagazine([]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "two", "is", "four"}))        //No
	fmt.Println(checkMagazine([]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}, []string{"ive", "got", "some", "coconuts"})) //No
}
