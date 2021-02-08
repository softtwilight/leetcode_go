package main

import (
	"fmt"
)

func numOfWays(nums []int) int {
	return int(numOfWaysHelper(nums) - 1)
}

func numOfWaysHelper(ns []int) int64 {

	if len(ns) <= 2 {
		return 1
	}
	left := make([]int, 0, len(ns))
	right := make([]int, 0, len(ns))
	for i := 1; i < len(ns); i++ {
		if ns[i] > ns[0] {
			left = append(left, ns[i])
		} else {
			right = append(right, ns[i])
		}
	}
	m := int64(1000000000 + 7)
	return comb(len(ns)-1, len(left)) * (numOfWaysHelper(left) % m) * (numOfWaysHelper(right) % m) % m

}

func comb(n, a int) int64 {
	result := int64(1)
	m := int64(1000000000 + 7)
	for i := 0; i < a; i++ {
		result *= int64(n)
		result %= m
		result /= int64(i + 1)
		n--
	}
	return result
}

func main() {
	fmt.Println("canPartition([]int{1,2,3,4,5,6,7})")
}
