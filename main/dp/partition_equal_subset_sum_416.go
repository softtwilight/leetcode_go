package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	memo := make(map[int]int)
	return helper1(sum/2, nums, len(nums)-1, memo)

}

func helper1(target int, nums []int, hi int, memo map[int]int) bool {
	if target == 0 {
		return true
	}
	if target < 0 || hi < 0 {
		return false
	}
	if v, ok := memo[target]; ok {
		if nums[v] == nums[hi] {
			return false
		}
	}
	for i := hi; i >= 0; i-- {
		if nums[hi] <= target {
			if helper1(target-nums[hi], nums, i-1, memo) {
				return true
			}
		}
	}
	memo[target] = hi
	return false
}

func main() {
	fmt.Println(canPartition([]int{2, 2, 1, 1}))
}
