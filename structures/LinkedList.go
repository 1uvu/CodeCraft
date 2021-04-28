package structures

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func ArrayConvertToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return head
}

func IsLinkedListEqual(h1, h2 *ListNode) bool {
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1 = h1.Next
		h2 = h2.Next
	}
	if h1 != nil || h2 != nil {
		return false
	}
	return true
}
