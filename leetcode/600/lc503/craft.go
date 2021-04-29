/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/next-greater-element-ii/"
@Tags    :   [单调栈]
---------------------------
@Idea:
 - Todo
 -
*/
package leetcode

func nextGreaterElements(nums []int) []int {
	var indexStack []int
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < 2*len(ans); i++ {
		for !empty(indexStack) && nums[top(indexStack)] < nums[i%len(ans)] {
			ans[top(indexStack)] = nums[i%len(ans)]
			indexStack = pop(indexStack)
		}
		indexStack = push(indexStack, i%len(ans))
	}
	return ans
}

func push(ns []int, n int) []int {
	ns = append(ns, n)
	return ns
}

func pop(ns []int) []int {
	return ns[:len(ns)-1]
}

func top(ns []int) int {
	return ns[len(ns)-1]
}

func empty(ns []int) bool {
	return len(ns) == 0
}
