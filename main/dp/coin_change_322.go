package main

import (
	"fmt"
)
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

// bottom up
func coinChange2(coins []int, amount int) int {

	memo := make([]int, amount + 1)
	for i := 1; i < len(memo); i++ {
		memo[i] = amount + 1
	}
	// sort to ascending order
	sort.Ints(coins)
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			memo[j] = min(memo[j], 1 + memo[j - coins[i]])
		}
	}
	if memo[amount] == amount + 1 {
		return -1
	}
	return memo[amount]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 5}
	re := coinChange2(nums, 14)
	fmt.Printf("the result = %v", re)
}
