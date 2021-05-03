/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/minimum-distance-to-the-target-element/"
@Tags    :   []
---------------------------
@Idea:
*/
package leetcode

import "math"

func getMinDistance(nums []int, target int, start int) int {
	ans := math.MaxInt64
	for i := range nums {
		if nums[i] == target && abs(start-i) < ans {
			ans = abs(start - i)
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
