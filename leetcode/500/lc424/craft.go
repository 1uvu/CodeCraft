/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/longest-repeating-character-replacement/"
@Tags    :   [双指针，滑动窗口]
---------------------------
@Idea:
 - 定义 cnt 记录当前窗口内各个字符的数量
 - maxLen = max(maxLen, cnt[s[right]-'A']) 根据右边界计算并记录历史最大长度
 - 当当前窗口无法满足凑成全部一样的字符 (right-left+1 > maxLen+k)，滑动左指针
 - 窗口是不停增加的，因此滑动完成后返回：len(s) - left
*/
package leetcode

func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}
	// 左指针 / 右指针
	left, right := 0, 0
	// 最长长度 / 当前窗口所有字符计数
	maxLen, cnt := 0, make([]int, 26)
	for right = 0; right < len(s); right++ {
		cnt[s[right]-'A']++
		maxLen = max(maxLen, cnt[s[right]-'A']) // 当前窗口最多的字符计数
		// 如果窗口无法满足凑成全部一样的字符，滑动左指针
		if right-left+1 > maxLen+k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
