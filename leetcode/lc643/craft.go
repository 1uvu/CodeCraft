/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/maximum-average-subarray-i/"
@Tags    :   [滑动窗口]
---------------------------
@Idea:
 - 一次移动一个元素，相加减即可
*/
package leetcode

func findMaxAverage(nums []int, k int) float64 {
	ms := 0
	for _, x := range nums[:k] {
		ms += x
	}
	_ms := ms
	for r:=k; r<len(nums); r++ {
		_ms = _ms - nums[r-k] + nums[r]
		if _ms > ms {
			ms = _ms
		}
	}

	return float64(ms)/float64(k)
}