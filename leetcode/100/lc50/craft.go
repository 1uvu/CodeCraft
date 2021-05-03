/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/powx-n/"
@Tags    :   [位运算]
---------------------------
@Idea:
 - 使用 幂次 * 2 代替 幂次 + 1
*/
package leetcode

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return binPow(x, n)
	} else {
		return 1.0 / binPow(x, -n)
	}
}

func binPow(x float64, n int) float64 {
	res := 1.0
	for n > 0 {
		if n&1 == 1 { // 奇数则在 x 自乘前， res 先乘一次 x （最终为 n == 1）
			res *= x
		}
		x *= x // 快速幂的精髓在这里，直接平方使得幂次乘以 2，比幂次加 1 快得多
		n >>= 1
	}
	return res
}
