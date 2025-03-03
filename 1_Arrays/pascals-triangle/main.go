package main

import "fmt"

// https://leetcode.com/problems/pascals-triangle/
func main() {
	input := 5
	res := generate(input)
	fmt.Printf("%v", res)
}

func generate(numRows int) [][]int {
	res := [][]int{}

	for i := range numRows {
		temp := []int{}

		for j := 0; j <= i; j++ {
			if j == 0 || j == len(res) {
				temp = append(temp, 1)
				continue
			}
			val := res[i-1][j-1] + res[i-1][j]
			temp = append(temp, val)
		}

		res = append(res, temp)
	}

	return res
}
