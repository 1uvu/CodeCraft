/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/brick-wall/"
@Tags    :   [Map，脑筋急转弯]
---------------------------
@Idea:
 - 利用 Map 统计重叠缝隙的个数
*/
package leetcode

func leastBricks(wall [][]int) int {
	M := make(map[int]int)
	for _, arr := range wall {
		i := 1
		for i < len(arr) {
			M[arr[i-1]]++
			arr[i] += arr[i-1]
			i++
		}
	}
	m := 0
	for _, v := range M {
		if v > m {
			m = v
		}
	}
	return len(wall) - m
}
