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

func main() {
	fmt.Println(countSubstrings("aaa"))
}