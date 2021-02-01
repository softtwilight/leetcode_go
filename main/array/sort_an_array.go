package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	sortArrayHelper(nums, 0, len(nums)-1)
	return nums
}

func sortArrayHelper(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	sortArrayHelper(nums, lo, p-1)
	sortArrayHelper(nums, p+1, hi)
}

func partition(nums []int, lo int, hi int) int {
	ra := rand.Intn(hi-lo+1) + lo
	lower := lo - 1
	exchange(nums, ra, hi)
	for i := lo; i < hi; i++ {
		if nums[i] <= nums[hi] {
			lower++
			exchange(nums, lower, i)
		}
	}
	exchange(nums, lower+1, hi)
	return lower + 1
}

func exchange(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0, 3}
	fmt.Printf("the result = %v", sortArray(nums))
}
