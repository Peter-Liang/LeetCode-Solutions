package problem114

import . "solutions/all/util"

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	right := root.Right
	flatten(root.Left)
	flatten(root.Right)
	root.Right = root.Left
	if root.Right == nil {
		root.Right = right
	} else {
		root.Left = nil
		cur := root.Right
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = right
	}
}

// @lc code=end
