/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/max-consecutive-ones-iii/"
@Tags    :   [双指针]
---------------------------
@Idea:
 - 将问题转化为：寻找和小于等于 K 的最长子数组（数组中元素非 0 即 1）
 - 先前移右指针，再前移左指针
*/
package leetcode

func longestOnes(A []int, K int) int {
	if K == len(A) {
		return K
	}
	sum := 0
	res := 0
	l, r := 0, 0
	for r < len(A) {
		for r < len(A) && sum <= K {
			sum += A[r]^1
			r++
		} // 右指针前移，此循环后 sum > K，也就是 sum = K+1

		if r == len(A) && (sum == K || l == 0) {
			// 处理特殊情况：遍历完成时前一个循环由于 r<len(A) 跳出循环可能使得 r 少加一次
			// 1、sum == K 意味着 r 少加了一次；
			// 2、l == 0 意味着整个数组都是 1，r 也少加一次。
			r++
		}
		if r-l-1 > res {
			// 记录最长子数组
			res = r-l-1
		}
		for l < r && sum > K {
			sum -= A[l]^1
			l++
		} // 左指针前移，此循环后 sum == K
	}
	return res
}