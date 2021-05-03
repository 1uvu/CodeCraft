/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/subsets/"
@Tags    :   [脑筋急转弯]
---------------------------
@Idea:
 - 以 [1,2,3] 为例,
	subsets[] 为 [[]]
	subsets[1]为 [[], [1]]
	subsets[1,2] 为 [[], [1], [2], [1,2]]
	......
	可以发现 (伪代码): subsets[1,2] = subsets[1] + [for sub in subset[1]: sub.append(2)]
*/
package leetcode

import "math"

func subsets(nums []int) [][]int {
	ans := make([][]int, 1, int(math.Pow(2, float64(len(nums))))+1)
	ans[0] = []int{}
	for _, x := range nums {
		for _, arr := range ans {
			a := make([]int, len(arr), len(arr)+1)
			copy(a, arr)
			a = append(a, x)
			ans = append(ans, a)
		}
	}
	return ans
}
