/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/"
@Tags    :   [递归]
---------------------------
@Idea:
 - 先使用 map 统计 ch 的频率，并遍历
 - 使用 val < k 的 key 分割 s，并遍历
 - 递归求解最大值
*/
package leetcode

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	for key, val := range freq(s) {
		if val < k {
			res := 0
			for _, _s := range split(s, key) {
				res = max(res, longestSubstring(_s, k))
			}
			return res
		}
	}
	return len(s)
}

func freq(s string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, ch := range s {
		_, ok := freqMap[ch]
		if ok {
			freqMap[ch]++
		} else {
			freqMap[ch] = 1
		}
	}
	return freqMap
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func split(s string, ch rune) []string {
	var sArray []string
	begin := 0
	for i, _ch := range s {
		if _ch == ch {
			sArray = append(sArray, s[begin:i])
			begin = i + 1
		}
	}
	sArray = append(sArray, s[begin:len(s)])
	return sArray
}
