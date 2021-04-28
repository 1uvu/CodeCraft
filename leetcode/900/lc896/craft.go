/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/monotonic-array/"
@Tags    :   []
---------------------------
@Idea:
 - 根据不同情况，判断单独递增、单调递减还是一直不变
*/
package leetcode

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	c := 1
	for c < len(A) && A[c]-A[c-1] == 0 {
		c++
	}
	if c == len(A) {
		return true
	}
	if A[c]-A[c-1] > 0 {
		return isMonAdd(A)
	} else if A[c]-A[c-1] < 0 {
		return isMonSub(A)
	} else {
		return isMonotonic(A[c:])
	}
}

func isMonAdd(A []int) bool {
	for i := 1; i < len(A); i++ {
		if A[i]-A[i-1] < 0 {
			return false
		}
	}
	return true
}

func isMonSub(A []int) bool {
	for i := 1; i < len(A); i++ {
		if A[i]-A[i-1] > 0 {
			return false
		}
	}
	return true
}
