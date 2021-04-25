/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/sum-of-unique-elements/"
@Tags    :   [Map]
---------------------------
@Idea:
 - freq 直接取巧根据数据范围建立长为 100 的数组，可以用 map 代替
 - 出现多次的数字 freq 标为 -1
*/
package leetcode

func sumOfUnique(nums []int) int {
	freq := make([]int, 100)
	for i:=0; i<len(nums); i++ {
		freq[nums[i]-1]++
		if freq[nums[i]-1] != 1 {
			if freq[nums[i]-1] != -1 {
				freq[nums[i]-1] = -1
			}

		}
	}
	sum := 0
	for i:=0; i<100; i++ {
		if freq[i] == 1 {
			sum += i+1
		}
	}
	return sum
}
