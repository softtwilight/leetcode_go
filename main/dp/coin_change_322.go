package main

import "fmt"
import "sort"

// https://leetcode.com/problems/coin-change/
// return the fewest number of coin make up the amount
func coinChange(coins []int, amount int) int {

	memo := make([]int, amount)
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	return dp(memo, coins, amount)

}

func dp(memo []int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount <= -1 {
		return -1
	}
	if memo[amount-1] != 0 {
		return memo[amount-1]
	}
	re := amount + 1
	for _, coin := range coins {
		if sub := dp(memo, coins, amount-coin); sub != -1 {
			if re > sub+1 {
				re = sub + 1
			}
		}
	}
	if re == amount+1 {
		re = -1
	}
	memo[amount-1] = re
	return re
}

func main() {
	nums := []int{1, 2, 5}
	re := coinChange(nums, 14)
	fmt.Printf("the result = %v", re)
}
