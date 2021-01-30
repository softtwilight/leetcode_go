package main

import (
	"fmt"
)

//https://leetcode.com/problems/counting-bits/

// 0 -> 1 ->
//10 -> 11 ->
//100 -> 101 -> 110 -> 111 ->
//1000 -> 1001 -> 1010 -> 1011 -> 1100 -> 1101 -> 1110 -> 1111 -> 10000
// 5 -> 1 + 0
func countBits(num int) []int {
	dp := make([]int, num+1)
	offset := 1
	for i := 1; i <= num; i++ {
		if i == offset*2 {
			dp[i] = 1
			offset <<= 1
		} else {
			dp[i] = dp[i-offset] + 1
		}
	}
	return dp
}

func main() {
	fmt.Println(countBits(16))
}
