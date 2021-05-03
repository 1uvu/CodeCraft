/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/"
@Tags    :   [栈]
---------------------------
@Idea:
 - 压栈，遇到相同的出栈
 - 返回最终栈里的元素
*/
package leetcode

type Stack []byte

func (stack Stack) String() string { return string(stack) }
func (stack Stack) Top() byte      { return stack[stack.Len()-1] }
func (stack Stack) Empty() bool    { return stack.Len() == 0 }
func (stack Stack) Len() int       { return len(stack) }

func removeDuplicates(S string) string {
	stack := make(Stack, 0, len(S))
	for i := range S {
		if stack.Empty() {
			stack = append(stack, S[i])
			continue
		}
		if !stack.Empty() && stack.Top() != S[i] {
			stack = append(stack, S[i])
		} else {
			stack = stack[:stack.Len()-1]
		}
	}
	return stack.String()
}
