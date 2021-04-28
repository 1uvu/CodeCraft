/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/reshape-the-matrix/"
@Tags    :   [矩阵]
---------------------------
@Idea:
 - 矩阵的索引映射：res[i][j] = nums[ (i*c+j) / len(nums[0]) ][ (i*c+j) % len(nums[0]) ]
*/
package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	/// 关键点：特殊情况处理、索引转换
	if r*c != len(nums)*len(nums[0]) {
		return nums
	}
	if len(nums) == r && len(nums[0]) == c {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			res[i][j] = nums[(i*c+j)/len(nums[0])][(i*c+j)%len(nums[0])]
		}
	}
	return res
}
