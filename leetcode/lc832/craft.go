/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/flipping-an-image/"
@Tags    :   [矩阵]
---------------------------
@Idea:
 - 交换并反转元素即可，时间复杂度：O(m*n/2)，空间复杂度：O(1)
*/
package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	n := len(A[0])
	for _, arr := range A {
		for j:=0; j<n/2; j++ {
			arr[j], arr[n-j-1] = 1 ^ arr[n-j-1], 1 ^ arr[j]
		}
		if n % 2 != 0 {
			arr[n/2] = 1 ^ arr[n/2]
		}
	}
	return A
}