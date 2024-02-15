package main

import (
	. "solutions/all/util"
)

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return square(getDepth(root.Left)-getDepth(root.Right)) <= 1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}

func square(num int) int {
	return num * num
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getDepth(root.Left), getDepth(root.Right))
}

// @lc code=end
