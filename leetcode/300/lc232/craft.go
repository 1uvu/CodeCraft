/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/implement-queue-using-stacks/"
@Tags    :   [栈，队列]
---------------------------
@Idea:
 - 使用双栈模拟队列
 - Push 时直接压入左栈
 - Pop 时首先从右栈弹出，如右栈为空，则将左栈清空，依次弹入右栈
 - Peek 先 Pop 再把元素压入右栈
*/
package leetcode

type MyQueue struct {
	leftStack  []int
	rightStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var queue MyQueue
	return queue
}

/** Push element x to the back of queue. */
func (queue *MyQueue) Push(x int) {
	queue.leftStack = append(queue.leftStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (queue *MyQueue) Pop() int {
	var x int
	if queue.rightEmpty() {
		for !queue.leftEmpty() {
			queue.rightStack = append(queue.rightStack, queue.leftStack[len(queue.leftStack)-1])
			queue.leftStack = queue.leftStack[:len(queue.leftStack)-1]
		}
		x = queue.rightStack[len(queue.rightStack)-1]
		queue.rightStack = queue.rightStack[:len(queue.rightStack)-1]
	} else {
		x = queue.rightStack[len(queue.rightStack)-1]
		queue.rightStack = queue.rightStack[:len(queue.rightStack)-1]
	}

	return x
}

/** Get the front element. */
func (queue *MyQueue) Peek() int {
	var x int
	x = queue.Pop()
	queue.rightStack = append(queue.rightStack, x)
	return x
}

/** Returns whether the queue is empty. */
func (queue *MyQueue) Empty() bool {
	return queue.leftEmpty() && queue.rightEmpty()
}

func (queue *MyQueue) leftEmpty() bool {
	return len(queue.leftStack) == 0
}

func (queue *MyQueue) rightEmpty() bool {
	return len(queue.rightStack) == 0
}
