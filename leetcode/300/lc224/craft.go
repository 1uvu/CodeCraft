/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/basic-calculator/submissions/"
@Tags    :   [栈]
---------------------------
@Idea:
 - 使用栈来处理括号
 - 使用 sign 来处理负数和减法
 - 循环读入处理连续的大整数
*/
package leetcode

var blank byte = ' '
var sub byte = '-'
var plus byte = '+'
var left byte = '('
var right byte = ')'

func calculate(s string) int {
	stack := []int{1}
	ans := 0
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		x := s[i]
		switch x {
		case blank:
			i++
		case plus: // 符号为 +
			sign = stack[len(stack)-1]
			i++
		case sub: // 符号为 -
			sign = -stack[len(stack)-1]
			i++
		case left: // 放入符号
			stack = append(stack, sign)
			i++
		case right: // 弹出括号内的结果
			stack = stack[:len(stack)-1]
			i++
		default: // 读入连续的数字
			num := 0
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			ans += sign * num
		}
	}
	return ans
}
