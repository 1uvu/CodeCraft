/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/toeplitz-matrix/"
@Tags    :   [矩阵]
---------------------------
@Idea:
 - 遍历 matrix[0][:n-1] 和 matrix[1:m-1][0]
 - 关键点：while 循环遍历对角元素的截止条件
*/
package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m-1; i++ {
		step := 1
		target := matrix[i][0]
		for step < n && step < m-i {
			next := matrix[i+step][step]
			if next != target {
				return false
			}
			step++
		}
	}
	for i := 1; i < n-1; i++ {
		step := 1
		target := matrix[0][i]
		for step < m && step < n-i {
			next := matrix[step][i+step]
			if next != target {
				return false
			}
			step++
		}
	}
	return true
}
