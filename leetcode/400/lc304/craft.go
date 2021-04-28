/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/range-sum-query-2d-immutable/"
@Tags    :   [矩阵]
---------------------------
@Idea:
 - 构建二维前缀和
 - 查找时要注意索引的计算
*/
package leetcode

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var nums NumMatrix
	if len(matrix) == 0 {
		return nums
	}
	sums := make([][]int, len(matrix))
	nums.sums = sums
	if len(matrix[0]) == 0 {
		return nums
	}
	for i := 0; i < len(matrix); i++ {
		arr := make([]int, len(matrix[0]))
		sums[i] = arr
	}

	// row = 0
	sums[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		sums[i][0] = matrix[i][0] + sums[i-1][0]
	}
	for j := 1; j < len(matrix[0]); j++ {
		sums[0][j] = matrix[0][j] + sums[0][j-1]
	}
	if len(matrix) == 1 || len(matrix[0]) == 1 {
		return nums
	}

	// row > 0
	for row := 1; row < len(matrix); row++ {
		arr := matrix[row]
		for col := 1; col < len(arr); col++ {
			sums[row][col] = matrix[row][col] + sums[row-1][col] + sums[row][col-1] - sums[row-1][col-1]
		}
	}
	return nums
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sums[row2][col2]
	} else if row1 == 0 && col1 != 0 {
		return this.sums[row2][col2] - this.sums[row2][col1-1]
	} else if row1 != 0 && col1 == 0 {
		return this.sums[row2][col2] - this.sums[row1-1][col2]
	} else {
		return this.sums[row2][col2] + this.sums[row1-1][col1-1] - this.sums[row1-1][col2] - this.sums[row2][col1-1]
	}
}
