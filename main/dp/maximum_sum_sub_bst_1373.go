package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int

func maxSumBST(root *TreeNode) int {
	result = 0
	maxSumBSTHelper(root)
	return result
}

func maxSumBSTHelper(root *TreeNode) (int, int, int, bool) {
	if root == nil {
		return 0, -10000000, 10000000, true
	}
	lsum, lmin, lmax, li := maxSumBSTHelper(root.Left)
	rsum, rmin, rmax, ri := maxSumBSTHelper(root.Right)
	isBst := li && ri && root.Val > lmax && root.Val < rmin
	if isBst {
		te := lsum + rsum + root.Val
		result = maxTwo(result, te)
		return te, lmin, rmax, true
	}
	return 0, lmin, rmax, false
}

func maxTwo(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println("hello")
}
