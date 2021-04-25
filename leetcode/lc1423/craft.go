/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/"
@Tags    :   [滑动窗口]
---------------------------
@Idea:
 - 找到一个和最小的长度为 len()-k 的滑动窗口
 - 总的和减去这个最小和即可
 - curSum = curSum+cardPoints[len(cardPoints)-k+l] - cardPoints[l]
*/
package leetcode

func maxScore(cardPoints []int, k int) int {

	minSum := sum(cardPoints[:len(cardPoints)-k])
	curSum := minSum
	for l:= 0; l<k; l++ {
		curSum = curSum+cardPoints[len(cardPoints)-k+l] - cardPoints[l]
		if curSum < minSum {
			minSum = curSum
		}
	}
	return sum(cardPoints)-minSum
}

func sum(ns []int) int {
	s:=0
	for _, x := range ns {
		s += x
	}
	return s
}