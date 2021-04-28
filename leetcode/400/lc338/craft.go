/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/counting-bits/"
@Tags    :   []
---------------------------
@Idea:
 -
 -
*/
package leetcode

func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		bits[i] = bits[i>>1] + (i & 1)
	}
	return bits
}
