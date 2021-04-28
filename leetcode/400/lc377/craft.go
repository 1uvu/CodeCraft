/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/combination-sum-iv/"
@Tags    :   [动态规划，背包]
---------------------------
@Idea:
 - 完全背包 + 有顺序组合问题（相同元素，不同顺序为不同组合）
 - dp 记录 1-target 的组合数目
 - if i >= x {
		dp[i] += dp[i - x]
	}
*/
package leetcode

func combinationSum4(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	i := 1
	for i <= target {
		for _, x := range nums {
			if i >= x {
				dp[i] += dp[i-x]
			}
		}
		i++
	}
	return dp[target]
}
