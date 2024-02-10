package problem109

import . "solutions/all/util"

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
	nodes := make([]*TreeNode, 0)
	cur := head
	for cur != nil {
		nodes = append(nodes, &TreeNode{Val: cur.Val})
		cur = cur.Next
	}

	mid := len(nodes) / 2
	root := nodes[mid]
	root.Left = nodesToBST(nodes[:mid])
	root.Right = nodesToBST(nodes[mid+1:])

	return root
}

func nodesToBST(nodes []*TreeNode) *TreeNode {
	switch size := len(nodes); size {
	case 0:
		return nil
	case 1:
		return nodes[0]
	case 2:
		nodes[1].Left = nodes[0]
		return nodes[1]
	}

	mid := len(nodes) / 2
	root := nodes[mid]
	root.Left = nodesToBST(nodes[:mid])
	root.Right = nodesToBST(nodes[mid+1:])
	return root
}

// @lc code=end
