/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/count-number-of-nice-subarrays/"
@Tags    :   [前缀和]
---------------------------
@Idea:
 - mp 记录奇数次数为 sum 的子数组个数，索引是 sum 前缀和，代表奇数个数，对应值是前缀和 sum 的个数（即到当前元素为止，数组中有多少个奇数为 sum 的子数组）
*/
package leetcode

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	var sum, ans int
	mp := make([]int, len(nums)+1)
	mp[0] = 1
	for i := range nums {
		if nums[i]%2 == 1 {
			sum++
		}
		mp[sum]++
		if sum-k >= 0 {
			// 加上与当前前缀和差值为 k 的前缀和的个数
			ans += mp[sum-k]
		}
	}
	fmt.Println(mp, sum)
	return ans
}
