package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	memo := make(map[int] int)
	return combHelper(nums, target, memo, 0)
}

func combHelper(nums []int, target int, memo map[int]int, i int) int {
	if target < 0 || i >= len(nums){
		return 0
	}
	if target == 0 {
		return 1
	}
	if v, ok := memo[target]; ok {
		return v
	}

	takeI := combHelper(nums, target - nums[i], memo, 0)
	notTakeI := combHelper(nums, target, memo, i + 1)
	memo[target] = takeI + notTakeI
	return memo[target]
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(combinationSum4(nums, 4))
}
