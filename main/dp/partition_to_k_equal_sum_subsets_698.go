package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	seen := make([]bool, len(nums))
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return dfs(seen, nums, k, sum/k, 0, len(nums)-1)

}

func dfs(seen []bool, nums []int, k int, target int, sum int, hi int) bool {
	if k == 0 {
		return true
	}
	if sum == target && dfs(seen, nums, k-1, target, 0, len(nums)-1) {
		return true
	}
	for i := hi; i >= 0; i-- {
		if !seen[i] && sum+nums[i] <= target {
			seen[i] = true
			if dfs(seen, nums, k, target, sum+nums[i], i-1) {
				return true
			}
			seen[i] = false
		}

	}
	return false
}

func main() {
	nums := []int{10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6}
	re := canPartitionKSubsets(nums, 3)
	fmt.Printf("result = %v", re)
}
