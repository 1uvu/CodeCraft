/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/daily-temperatures/"
@Tags    :   [栈，动态规划]
---------------------------
@Idea:
 - 使用栈存储温度的坐标
 - 使用 ans[stack[top]] = i - stack[top] 直接求得天数差
*/
package leetcode

func dailyTemperatures(T []int) []int {
	if len(T) == 1 {
		return []int{0}
	}
	ans := make([]int, len(T))
	stack := make([]int, len(T))
	top := 0
	for i := 0; i < len(T); i++ {
		for top > 0 && T[i] > T[stack[top]] {
			ans[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}

	return ans
}
