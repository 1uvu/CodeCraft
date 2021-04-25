/* -*- encoding: utf-8 -*- */
/*
@File    :   craft.go
@URL     :   "https://leetcode-cn.com/problems/increasing-order-search-tree/"
@Tags    :   [二叉搜索树]
---------------------------
@Idea:
 - 中序遍历得到数组
 - 定义哑节点，生成新树
*/
package leetcode

import "github.com/1uvu/codecraft/structures"

type TreeNode = structures.BNode

func increasingBST(root *TreeNode) *TreeNode {
	a := structures.InorderTraverseBinaryTree(root)
	tummy := &TreeNode{-1, nil, nil}
	node := &TreeNode{a[0].(int), nil, nil}
	tummy.Right = node
	for i:=1; i<len(a); i++ {
		node.Left = nil
		node.Right = &TreeNode{a[i].(int), nil, nil}
		node = node.Right
	}
	return tummy.Right
}
