package _go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	result := rotateString("testing", 3, 1)
	fmt.Println("Expected: 'ngtesti' - Result: ", result)
}
