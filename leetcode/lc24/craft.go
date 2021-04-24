/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/swap-nodes-in-pairs/"
@Tags    :   [链表]
---------------------------
@Idea:
 - 递归，两两交换
 - 迭代，画图立马就理解
*/
package leetcode

import (
	"github.com/1uvu/codecraft/structures"
)

type ListNode = structures.ListNode

// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 迭代
func swapPairs_(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n2 := cur.Next.Next
		cur.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		cur = n1
	}
	return dummyHead.Next
}