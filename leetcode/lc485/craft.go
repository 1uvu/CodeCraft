/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/max-consecutive-ones/submissions/"
@Tags    :   []
---------------------------
@Idea:
 - 一次遍历
*/
package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	m := 0
	_m := 0
	for _, x := range nums {
		if x == 0 {
			if _m > m {
				m = _m
			}
			_m = -1
		}
		_m++
	}
	if _m > m {
		m = _m
	}
	return m
}