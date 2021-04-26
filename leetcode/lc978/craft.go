/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/longest-turbulent-subarray/"
@Tags    :   [滑动窗口，动态规划]
---------------------------
@Idea:
 - 问题转化：通过 compare 函数将 arr 转化为只有 1，-1 和 0 的形式
 - -1，1 或 1，-1 时 arr[i] + arr[i-1] = 0 代表湍流
 - 即求相加之和为 [-1, 1] 的最长子序列
*/
package leetcode

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	flag := true  // 用来判断全部相等的情况
	m, _m := 0, 0
	for r := 0; r < len(arr)-1; r++ {
		arr[r] = compare(arr[r], arr[r+1])
		if arr[r] != 0 {
			flag = false
		}
		if r > 0 && arr[r] != 0 && arr[r] + arr[r-1] == 0 {
			_m++
		} else {
			_m = 0
		}
		if _m > m {
			m = _m
		}
	}
	if flag {
		return 1
	}
	return m + 2
}

func compare(a int, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}
