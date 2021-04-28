/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/get-equal-substrings-within-budget/"
@Tags    :   [滑动窗口]
---------------------------
@Idea:
 - 问题转化：寻找 cost 数组中和加起来小于等于 maxCost 的最大子序列`长度`
*/
package leetcode

func equalSubstring(s string, t string, maxCost int) int {
	maxLen := 0
	curCost := 0
	l := 0

	for r := 0; r < len(s); r++ {
		curCost += abs(int(t[r]) - int(s[r]))
		if curCost <= maxCost {
			if r+1 < len(s) {
				if curCost+abs(int(t[r+1])-int(s[r+1])) > maxCost && r-l+1 >= maxLen {
					maxLen = r - l + 1
				}
			} else {
				if r-l+1 >= maxLen {
					maxLen = r - l + 1
				}
			}

		} else {
			curCost = curCost - abs(int(t[l])-int(s[l]))
			l++
		}
	}

	return maxLen
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
