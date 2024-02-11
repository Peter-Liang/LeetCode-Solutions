package problem109

import (
	. "solutions/all/util"
)

/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	size := 1
	nodes := &TreeNode{Val: head.Val}
	next := nodes
	cur := head.Next
	for cur != nil {
		next.Right = &TreeNode{Val: cur.Val}
		next = next.Right
		cur = cur.Next
		size++
	}

	return nodesToBST(nodes, size)
}

func nodesToBST(node *TreeNode, size int) *TreeNode {
	switch size {
	case 0:
		return nil
	case 1:
		node.Left = nil
		node.Right = nil
		return node
	case 2:
		root := node.Right
		root.Left = node
		root.Right = nil
		node.Left = nil
		node.Right = nil
		return root
	}

	mid := size / 2
	var root *TreeNode
	cur := node
	for i := 0; i < mid; i++ {
		cur = cur.Right
	}
	root = cur
	root.Left = nodesToBST(node, mid)
	root.Right = nodesToBST(cur.Right, size-mid-1)
	return root
}

// @lc code=end
