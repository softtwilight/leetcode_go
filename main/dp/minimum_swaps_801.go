package main

import "fmt"

func minSwap(A []int, B []int) int {

	memo := make([][]int, len(A))
	for i := range memo {
		memo[i] = make([]int, 2)
	}
	return helperMinSwap(A, B, 1, 0, memo)

}

func helperMinSwap(A []int, B []int, i int, swaped int, memo [][]int) int {

	if i >= len(A) {
		return 0
	}
	if memo[i][swaped] > 0 {
		return memo[i][swaped] - 1
	}
	lastA := 0
	lastB := 0
	if swaped == 0 {
		lastA = A[i-1]
		lastB = B[i-1]
	} else {
		lastA = B[i-1]
		lastB = A[i-1]
	}
	result := 2001
	if A[i] > lastB && B[i] > lastA {
		result = 1 + helperMinSwap(A, B, i+1, 1, memo)
	}
	if A[i] > lastA && B[i] > lastB {
		notSwap := helperMinSwap(A, B, i+1, 0, memo)
		if notSwap < result {
			result = notSwap
		}
	}
	memo[i][swaped] = result + 1
	return result
}

func main() {
	A := []int{2, 1, 6, 7, 8, 13, 15, 11, 18, 13, 20, 24, 17, 28, 22, 23, 36, 37, 39, 34, 43, 38, 48, 41, 46, 48, 49, 50, 56, 55, 59, 60, 62, 64, 66, 75, 69, 70, 71, 74, 87, 78, 95, 97, 81, 99, 85, 101, 90, 91, 93, 95, 107, 109, 101, 111, 106, 114, 115, 117, 118, 115, 121, 122, 123, 124, 125, 126, 134, 131, 133, 136, 142, 149, 151, 152, 145, 156, 158, 150, 162, 159, 161, 165, 169, 170, 169, 174, 172, 176, 177, 181, 183, 192, 186, 188, 189, 196, 198, 200}
	B := []int{0, 4, 10, 11, 12, 9, 10, 16, 12, 19, 15, 16, 25, 20, 33, 34, 27, 29, 32, 40, 35, 45, 40, 50, 51, 52, 53, 55, 52, 58, 58, 61, 62, 66, 71, 68, 78, 81, 83, 84, 75, 91, 79, 80, 98, 83, 100, 89, 102, 103, 105, 106, 96, 98, 110, 105, 113, 109, 110, 111, 112, 120, 116, 118, 126, 130, 131, 133, 129, 137, 138, 140, 137, 138, 140, 142, 154, 147, 149, 159, 152, 163, 164, 163, 166, 168, 171, 170, 175, 176, 177, 181, 186, 184, 193, 194, 195, 190, 195, 200}
	fmt.Println(minSwap(A, B))
}
