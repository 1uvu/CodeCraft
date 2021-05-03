/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/degree-of-an-array/submissions/"
@Tags    :   [Map]
---------------------------
@Idea:
 - leftMap 存储元素第一次出现的位置
 - rightMap 存储最后一次出现的位置
 - freqMap 存储元素的频率， maxFreq 记录最大频率
 - 先遍历一次求得上面的值
 - 最后根据 maxFreq 求得与其频率相等的元素的最短长度
*/
package leetcode

func findShortestSubArray(nums []int) int {
	freqMap := make(map[int]int)
	leftMap := make(map[int]int)
	rightMap := make(map[int]int)
	maxFreq := 0
	for i, x := range nums {
		if _, ok := freqMap[x]; ok {
			freqMap[x]++
		} else {
			freqMap[x] = 1
			leftMap[x] = i
			rightMap[x] = i
		}
		if freqMap[x] >= maxFreq {
			maxFreq = freqMap[x]
			rightMap[x] = i
		}
	}
	minLen := len(nums)
	for key := range freqMap {
		if freqMap[key] == maxFreq && rightMap[key]-leftMap[key]+1 < minLen {
			minLen = rightMap[key] - leftMap[key] + 1
		}
	}
	return minLen
}
