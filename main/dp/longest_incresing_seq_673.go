package main

import "fmt"

//https://leetcode.com/problems/number-of-longest-increasing-subsequence/
func findNumberOfLIS(nums []int) int {
	maxLen := make([]int, len(nums))

	// count the #ways to reach i by longest increasing sequence.
	cnt := make([]int, len(nums))
	longest := 0
	re := 0
	for i := 0; i < len(nums); i++ {
		cnt[i], maxLen[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if maxLen[i] == maxLen[j] + 1 {
					cnt[i] += cnt[j]
				}
				if maxLen[i] < maxLen[j] + 1 {
					maxLen[i] = maxLen[j] + 1
					cnt[i] = cnt[j]
				}
 			}
		}
		if longest == maxLen[i] {
			re += cnt[i]
		}
		if longest < maxLen[i]{
			longest = maxLen[i]
			re = cnt[i]
		}
	}
	return re
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	nums := []int{1,3,5,4,7, 7}
	fmt.Println(findNumberOfLIS(nums))
}