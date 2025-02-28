package main

import "fmt"

// https://leetcode.com/problems/set-matrix-zeroes/
func main() {
	input := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(input)
}

// TC|SC : O(n*m)|O(1)
// if a '0' appears while traversing the matrix [i][j], we will mark the first element of that row([0][j]) and column([i][0]) as '0'.
// we need to handle the edge case where the first row or column might be having '0' to begin with, this is done using the firstRow, firstCol flags.
// after traversing throughout the array, 
// we traverse through the first row and columns and if they are '0' we will set all the elements of that row and column to '0'
func setZeroes(matrix [][]int) {

	firstRow, firstCol := false, false
	n, m := len(matrix), len(matrix[0])
	for i := range n {
		for j := range m {
			val := matrix[i][j]
			if val == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstCol = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i<n; i++ {
		val := matrix[i][0]
		if val == 0 {
			for j := range m {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i<m; i++ {
		val := matrix[0][i]
		if val == 0 {
			for j := range n {
				matrix[j][i] = 0
			}
		}
	}

	if firstRow {
		for i := range m {
			matrix[0][i] = 0
		}
	}

	if firstCol {
		for i := range n {
			matrix[i][0] = 0
		}
	}

	fmt.Println(matrix)
}
