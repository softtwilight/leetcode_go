package main

import (
	"fmt"
)

const mod = int(1e9 + 7)

func numOfWays(nums []int) int {
	return (numOfWaysHelper(nums) - 1 + mod) % mod
}

func numOfWaysHelper(ns []int) int {

	if len(ns) <= 2 {
		return 1
	}
	left := []int{}
	right := []int{}
	for i := 1; i < len(ns); i++ {
		if ns[i] > ns[0] {
			left = append(left, ns[i])
		} else {
			right = append(right, ns[i])
		}
	}
	return comb(len(ns)-1, len(left)) * (numOfWaysHelper(left) % mod) * (numOfWaysHelper(right) % mod)

}

func comb(n, a int) int {

	result := 1
	for i := 0; i < a; i++ {
		result *= n
		result %= mod
		result /= (i + 1)
		n--
	}
	return result
}

func main() {
	fmt.Println("canPartition([]int{1,2,3,4,5,6,7})")
}
