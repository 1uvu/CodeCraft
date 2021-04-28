/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/permutation-in-string/"
@Tags    :   [滑动窗口]
---------------------------
@Idea:
 - 使用两个 26 长的数组记录字母数目
 - 对于不同长度的情况，每额外添加一个字母就调用一次 compare
*/
package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	r := len(s1)
	C1, C2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s1); i++ {
		C1[s1[i]-'a']++
		C2[s2[:r][i]-'a']++
	}
	if compare(C1, C2) {
		return true
	}
	for r < len(s2) {
		C2[s2[r]-'a']++
		C2[s2[r-len(s1)]-'a']--
		if compare(C1, C2) {
			return true
		}
		r++
	}

	return false
}

func compare(C1 []int, C2 []int) bool {
	for i := 0; i < len(C1); i++ {
		if C1[i] != C2[i] {
			return false
		}
	}

	return true
}
