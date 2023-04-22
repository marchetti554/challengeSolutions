package _go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println("maxProfit() --> ", result)
	result = maxProfitImproved([]int{7, 1, 5, 3, 6, 4})
	fmt.Println("maxProfitImproved() --> ", result)
}
