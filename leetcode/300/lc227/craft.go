/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/basic-calculator-ii/"
@Tags    :   [栈]
---------------------------
@Idea:
 - 使用栈保存中间结果
 - 使用 sign 来处理负数和减法
 - 循环读入处理连续的大整数
 - 无括号，所以遇到 * / 直接计算
*/
package leetcode

func calculate(s string) (ans int) {
	var stack []int
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}
