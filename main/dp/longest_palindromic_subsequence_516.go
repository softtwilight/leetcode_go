package main

import "fmt"

func longestPalindromeSubseq(s string) int {

	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}
	return longestPalindromeSubseqHelper(s, 0, len(s)-1, memo)
}

func longestPalindromeSubseqHelper(s string, lo int, hi int, memo [][]int) int {
	if lo == hi {
		return 1
	}
	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}
	result := 0
	if s[lo] == s[hi] {
		if lo+1 == hi {
			return 2
		}
		result = 2 + longestPalindromeSubseqHelper(s, lo+1, hi-1, memo)
	} else {
		result = longestPalindromeSubseqHelper(s, lo+1, hi, memo)
		if v := longestPalindromeSubseqHelper(s, lo, hi-1, memo); v > result {
			result = v
		}
	}
	memo[lo][hi] = result
	return result
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
