/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/maximum-score-from-removing-stones/"
@Tags    :   [脑筋急转弯]
---------------------------
@Idea:
 - 找规律
*/
package leetcode

func maximumScore(a int, b int, c int) int {
	m := max(a, b, c)
	if a+b+c <= 2*m {
		return (a + b + c) - m
	} else {
		return (a + b + c) >> 1
	}
}

func max(a, b, c int) int {
	if b < a && c < a {
		return a
	}
	if a < b && c < b {
		return b
	}
	return c
}
