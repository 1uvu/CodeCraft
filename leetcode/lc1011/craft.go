/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/"
@Tags    :   [二分查找]
---------------------------
@Idea:
 - 找到一个`最小的`装载容量 <=> 下限为 weights 中的最大值，上限为 weights 的总和
 - 根据这个界限，二分查找到`第一个`可以在 D 天内完成装载的容量，就是最小的
 - QA：为什么将上限根据 D 设为平均值不可以？因为分布可能是不均匀的，只能将上限设为总和。
*/
package leetcode

import "sort"

func shipWithinDays(weights []int, D int) int {
	// 确定二分查找左右边界
	// left: max
	// right: sum
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	// 返回二分查找成功的第一个结果
	return left + sort.Search(right - left, func(x int) bool {
		// x 是 [0, right - left] 内的一个数字
		// x += left 代表着当前设置的运载容量
		// 判断传入的 x + left 是否能够在 D 天内运输完
		//
		x += left
		day := 1 // 需要运送的天数
		sum := 0 // 当前这一天已经运送的包裹重量之和
		for _, w := range weights {
			if sum + w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}