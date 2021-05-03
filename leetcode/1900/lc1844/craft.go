/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/replace-all-digits-with-characters/"
@Tags    :   []
---------------------------
@Idea:
*/
package leetcode

func replaceDigits(s string) string {
	ans := make([]byte, len(s))
	for i := range s {
		ch := s[i]
		if i%2 == 1 {
			ans[i] = s[i-1] + ch - '0'
		} else {
			ans[i] = ch
		}
	}
	return string(ans)
}
