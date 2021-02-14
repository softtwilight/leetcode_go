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
		return 0, 0, 0, true
	}
	lsum, lmin, lmax, li := maxSumBSTHelper(root.Left)
	rsum, rmin, rmax, ri := maxSumBSTHelper(root.Right)
    isBst := li && ri && (lsum == 0 || root.Val > lmax)  && (rsum == 0 || root.Val < rmin)
    if lsum == 0 {
        lmin = root.Val
        lmax = root.Val
    }
    if rsum == 0 {
        rmin = root.Val
        rmax = root.Val
    }
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
