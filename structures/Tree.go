package structures

import (
	"github.com/1uvu/codecraft/utils"
)

type BNode struct {
	Val   interface{}
	Left  *BNode
	Right *BNode
}

func ArrayConvertToBinaryTree(nums []int) *BNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &BNode{nums[0], nil, nil}
	queue := make([]*BNode, 1, n*2) // 叶子节点的子节点为 NULL
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != utils.NULL {
			node.Left = &BNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != utils.NULL {
			node.Right = &BNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func InorderTraverseBinaryTree(root *BNode) []interface{} {
	var ans []interface{}
	var traverse func(node *BNode)
	traverse = func(root *BNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		ans = append(ans, root.Val)
		traverse(root.Right)
	}
	traverse(root)
	return ans
}

func IsBinaryTreeInOrderEqual(r1, r2 *BNode) bool {
	//  判断两棵二叉树是否具有中序遍历结果
	a1, a2 := InorderTraverseBinaryTree(r1), InorderTraverseBinaryTree(r2)
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func IsBinaryTreeEqual(r1, r2 *BNode) bool {
	//  判断两棵二叉树是否具有相同的结构和数值
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.Val == r2.Val && IsBinaryTreeEqual(r1.Left, r2.Left) && IsBinaryTreeEqual(r1.Right, r2.Right)
}
