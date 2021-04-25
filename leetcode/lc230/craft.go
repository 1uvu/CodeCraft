/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/"
@Tags    :   [二叉搜索树]
---------------------------
@Idea:
 - 定义 order 代表一次遍历
 - k--
*/
package leetcode

import (
	"github.com/1uvu/codecraft/structures"
)

type TreeNode = structures.BNode

func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		order(root.Left)
		k--
		if k == 0 {
			ans = root.Val.(int)
			return
		}
		order(root.Right)
	}
	order(root)
	return ans
}

