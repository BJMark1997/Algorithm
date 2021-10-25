package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if n*m != r*c {
		return nums
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for i := 0; i < n*m; i++ {
		ans[i/c][i%c] = nums[i/m][i%m]
	}
	return ans
}

func main() {
	input := [][]int{{1, 2}, {3, 4}, {5, 6}}
	reshape := matrixReshape(input, 2, 3)
	fmt.Print(reshape)
}
