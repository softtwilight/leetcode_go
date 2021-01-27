package main

import (
	"fmt"
)

//https://leetcode.com/problems/partition-equal-subset-sum/

func canPartition(nums []int) bool {
	dp := make(map[int]bool)

	targetSum := 0
	for _, val:= range nums {
		targetSum += val
	}

	if targetSum %2 !=0 {
		return false
	}

	//sumSet := []int {}
	return canPartitionHelper(targetSum/2, 0, nums, dp)
}
func canPartitionHelper( target int,  i int, nums[]int, memo map[int] bool) bool {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == 0 {
		return true
	}
	if i == len(nums){
		return false
	}
	if nums[i] <= target {
		if canPartitionHelper(target-nums[i], i+1, nums, memo) == true {
			memo[target] = true
			return true
		}
	}
	memo[target] = canPartitionHelper(target, i+1, nums, memo)
	return memo[target]
}

func canPartition2(nums []int) bool {
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
	fmt.Println(canPartition([]int{1,2,3,4,5,6,7}))
	fmt.Println(canPartition3([]int{1,2,3,4,5,6,7}))
}
