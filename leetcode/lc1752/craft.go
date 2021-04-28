/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/check-if-array-is-sorted-and-rotated/"
@Tags    :   []
---------------------------
@Idea:
 - 轮转之后，找到一个中间的 index，index 前面为递减，后面为递增
 -
*/
package leetcode

func check(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	index := 0
	for index < len(nums) {
		if index == 0 {
			if nums[index] > nums[index+1] {
				index++
				break
			}
		} else if index > 0 && index < len(nums)-1 {
			if nums[index] > nums[index+1] || nums[index] < nums[index-1] {
				index++
				break
			}
		} else {
			if nums[index] < nums[index-1] {
				break
			}
		}
		index++
	}
	return judge(append(nums[index:], nums[:index]...))
}

func judge(nums []int) bool {
	i := 0
	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			return false
		}
		i++
	}
	return true
}
