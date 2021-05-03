/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/palindrome-partitioning/"
@Tags    :   [回溯]
---------------------------
@Idea:
 - backtrack 模板
 - 1 return 条件：到达时 append path 的副本
 - 2 循环递归
 - 3 递归条件：判断条件是否进行回溯，先 append 下一步，回溯，再 pop
*/
package leetcode

func partition(s string) [][]string {
	var ans [][]string
	var path []string
	var backtrack func(s string, path []string)
	backtrack = func(s string, path []string) {
		if len(s) == 0 {
			copyPath := make([]string, len(path))
			copy(copyPath, path)
			ans = append(ans, copyPath)
			return
		}
		for i := 1; i < len(s)+1; i++ {
			if isHuiWen(s[:i]) {
				path = append(path, s[:i])
				backtrack(s[i:], path)
				path = path[:len(path)-1]
			}
		}
		return
	}
	backtrack(s, path)
	return ans
}

func isHuiWen(s string) bool {
	for i := len(s) / 2; i >= 0; i-- {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
