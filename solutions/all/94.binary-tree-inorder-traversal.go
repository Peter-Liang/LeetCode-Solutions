package problem_94

// import "playground/main/utils/TreeNode"

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		if node == nil {
			nodes = nodes[:len(nodes)-1]
		} else if node.Left != nil {
			// handle left
			nodes = append(nodes, node.Left)
			node.Left = nil
		} else {
			// handle current and right
			nodes = nodes[:len(nodes)-1]
			result = append(result, node.Val)
			nodes = append(nodes, node.Right)
		}
	}

	return result
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
