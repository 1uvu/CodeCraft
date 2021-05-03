/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/single-number-ii/"
@Tags    :   [位运算，Map]
---------------------------
@Idea:
 - 相同数字的相应二进制位 %3 == 0
 - 因此统计那些不为 0 的在对应位加起来就是最终的结果
*/
package leetcode

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1 // 统计 nums 中二进制 i 位为 1 的个数
		}
		if total%3 > 0 { // 如果 i 位的 1 的个数不是 3 的倍数，则说明这里多了一个，则加到 ans 里
			ans |= 1 << i
		}
	}
	return int(ans)
}
