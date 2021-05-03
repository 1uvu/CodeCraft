/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/valid-anagram/"
@Tags    :   [Map]
---------------------------
@Idea:
 - 使用长度为 26 的数组来对字母进行计数：同时出现为 0 否则为 1/-1
 - 最后遍历一次检查次数是否相同
*/
package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}

	var countArray [26]int

	for i := 0; i < len(s); i++ {
		countArray[s[i]-'a']++
		countArray[t[i]-'a']--
	}

	for j := 0; j < 26; j++ {
		if countArray[j] != 0 {
			return false
		}
	}

	return true
}
