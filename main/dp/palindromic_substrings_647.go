package main

import (
	"fmt"
)

//https://leetcode.com/problems/palindromic-substrings/

// brute force version
// check all the subString
func countSubstrings(s string) int {
	re := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++{
			if isPalindromic(s[i : j]) {
				re++
			}
		}
	}
	return re
}

func isPalindromic(s string) bool {
	lo := 0
	hi := len(s) - 1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}


// a more efficient version, expand from mid
// the change the mid index
// from one index or two
func countSubstrings2(s string) int {
	sl := len(s)
	count := 0
	for mid := 0; mid < 2 * sl - 1; mid++ {
		lo := mid / 2
		hi := lo + mid % 2
		for lo >= 0 && hi < sl && s[lo] == s[hi] {
			count++
			lo--
			hi++
		}
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("abccb"))
	fmt.Println(countSubstrings2("abccb"))
}