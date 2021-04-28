/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/transpose-matrix/"
@Tags    :   [矩阵]
---------------------------
@Idea:
*/
package leetcode

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, n)
	for j := 0; j < n; j++ {
		arr := make([]int, m)
		for i := 0; i < m; i++ {
			arr[i] = matrix[i][j]
		}
		res[j] = arr
	}
	return res
}
