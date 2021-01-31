package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	return maxProfitHelper(prices, 0, dp)
}

// for n every function call, n sub problem
// time complexity O(n^2)
// space complexity O(n)
func maxProfitHelper(prices []int, start int, dp []int) int {
	for start < len(prices)-1 && prices[start] > prices[start+1] {
		start++
	}
	if start >= len(prices) {
		return 0
	}
	if dp[start] != 0 {
		return dp[start] - 1
	}

	buyAtStart := 0
	for i := start + 1; i < len(prices); i++ {
		if prices[i] > prices[start] {
			re := maxProfitHelper(prices, i+2, dp) + prices[i] - prices[start]
			if re > buyAtStart {
				buyAtStart = re
			}
		}
	}
	passStart := maxProfitHelper(prices, start+1, dp)
	if passStart > buyAtStart {
		buyAtStart = passStart
	}
	dp[start] = buyAtStart + 1
	return buyAtStart
}

func main() {
	nums := []int{1, 2, 5, 0, 2, 5}
	fmt.Println(maxProfit(nums))
}
