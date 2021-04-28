/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/grumpy-bookstore-owner/"
@Tags    :   [滑动窗口]
---------------------------
@Idea:
 - 直接将 grumpy 与 customers 相乘存入 grumpy
 - 计算连续子数组的最大和：_m = _m + grumpy[ptr] - grumpy[ptr-X]
*/
package leetcode

func maxSatisfied(customers []int, grumpy []int, X int) int {
	r := 0
	for i, x := range customers {
		grumpy[i] *= x
		if grumpy[i] == 0 {
			r += x
		}
	}
	m := 0
	for _, x := range grumpy[:X] {
		m += x
	}
	_m := m
	for ptr:=X; ptr<len(grumpy); ptr++ {
		_m = _m + grumpy[ptr] - grumpy[ptr-X]
		if _m > m {
			m = _m
		}
	}
	return r+m
}