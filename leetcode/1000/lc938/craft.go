/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/range-sum-of-bst/"
@Tags    :   [二叉搜索树，DFS]
---------------------------
@Idea:
 - 中序遍历二叉搜索树，当 Val 落在范围内是 res +=
*/
package leetcode

import (
	"github.com/1uvu/codecraft/structures"
)

type TreeNode = structures.BNode

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	res := 0
	var DFS func(*TreeNode)
	DFS = func(root *TreeNode) {
		if root.Val.(int) >= low && root.Val.(int) <= high {
			res += root.Val.(int)
		}
		if root.Right != nil {
			DFS(root.Right)
		}
		if root.Left != nil {
			DFS(root.Left)
		}
		if root.Left == nil && root.Right == nil {
			return
		}
	}
	DFS(root)
	return res
}
