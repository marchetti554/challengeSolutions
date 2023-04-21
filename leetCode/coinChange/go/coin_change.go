package _go

// Incomplete
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 1 {
		if coins[0] == amount {
			return 1
		}
		return -1
	}
	//var totalCoinAmount int
	for i := len(coins) - 1; i != 0; i-- {
		//if coins[i]
		//	totalCoinAmount += amount / coins[i]
		//rest := amount % coins[i]
		//i

	}

	return -1
}

//coins = [1,2,5], amount = 11
//coins = [2,5,9], amount = 30 ---> output -> 3 de 9 + 1 de 2 + 1 de 1, total=5

//1. We divide our amount for out highest coin (30/9)
//2. That would be 3, with a rest of 3 too.
//3. Now, we check if we can sum up to 3 with the rest of coins.
