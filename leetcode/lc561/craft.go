/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/array-partition-i/"
@Tags    :   []
---------------------------
@Idea:
 - 排序后遍历
*/
package leetcode

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)/2; i++ {
		ans += nums[2*i]
	}
	return ans
}
