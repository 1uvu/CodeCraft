/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/fair-candy-swap/"
@Tags    :   [双指针，数学问题]
---------------------------
@Idea:
 - 求出 avg = (SumA - SumB)/2，A 和 B 的交换目标之差等于 avg
 - 即：A[Ai] - B[Bi] == avg
*/
package leetcode

import "sort"

func fairCandySwap(A []int, B []int) []int {
	avg := 0
	for i := 0; i < len(A); i++ {
		avg += A[i]
	}
	for j := 0; j < len(B); j++ {
		avg -= B[j]
	}
	avg /= 2
	sort.Ints(A)
	sort.Ints(B)
	Ai := 0
	Bi := 0
	for A[Ai]-B[Bi] != avg {
		if A[Ai]-B[Bi] > avg {
			Bi++
		} else if A[Ai]-B[Bi] < avg {
			Ai++
		} else {
			break
		}
	}
	return []int{A[Ai], B[Bi]}
}
