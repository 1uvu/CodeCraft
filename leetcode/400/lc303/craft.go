/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/range-sum-query-immutable/"
@Tags    :   [前缀和]
---------------------------
@Idea:
 - 初始化构造前缀和
*/
package leetcode

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	var obj NumArray
	if len(nums) == 0 {
		return obj
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	obj.sums = sums
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return obj
}

func (obj *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return obj.sums[j]
	} else {
		return obj.sums[j] - obj.sums[i-1]
	}
}
