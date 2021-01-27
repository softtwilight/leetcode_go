package main

import (
	"fmt"
)

//https://leetcode.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum % 2 != 0 {
		return false
	}
	sum = sum / 2
	memo := make([]bool, sum + 1)
	memo[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j - nums[i]]
		}
	}
	return memo[sum]
}

//
//	memo := make(map[int]int)
//	return helper1(sum/2, nums, len(nums)-1, memo)
//
//}

//func helper1(target int, nums []int, hi int, memo map[int]int) bool {
//	if target == 0 {
//		return true
//	}
//	if target < 0 || hi < 0 {
//		return false
//	}
//	if v, ok := memo[target]; ok {
//		if nums[v] == nums[hi] {
//			return false
//		}
//	}
//	for i := hi; i >= 0; i-- {
//		if nums[hi] <= target {
//			if helper1(target-nums[hi], nums, i-1, memo) {
//				return true
//			}
//		}
//	}
//	memo[target] = hi
//	return false
//}

func main() {
	fmt.Println(canPartition([]int{2, 2, 1, 1}))
}
