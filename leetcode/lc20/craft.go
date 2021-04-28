/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/valid-parentheses/"
@Tags    :   [栈]
---------------------------
@Idea:
 - 使用数组模拟栈，直接初始化为 len(s)
 - 使用 p 记录数组中当前栈顶的位置
 - 匹配则出栈，不匹配则入栈
*/
package leetcode

func isValid(s string) bool {
	stack := make([]byte, len(s))
	top := -1
	for i := 0; i < len(s); i++ {
		if top == -1 {
			top++
			stack[top] = s[i]
			continue
		}
		switch stack[top] {
		case '{':
			if s[i] == '}' { // 匹配则将栈顶指针指向的元素弹出
				top--
			} else { // 不匹配则将当前元素入栈
				top++
				stack[top] = s[i]
			}

		case '(':
			if s[i] == ')' {
				top--
			} else {
				top++
				stack[top] = s[i]
			}

		case '[':
			if s[i] == ']' {
				top--
			} else {
				top++
				stack[top] = s[i]
			}
		}
	}
	// 栈空则全部匹配，返回 true
	if top == -1 {
		return true
	} else {
		return false
	}
}
