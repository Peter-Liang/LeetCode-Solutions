package problem99

import (
	"slices"
)

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	sortedVals := treeToSlice(root)
	originVals := treeToSlice(root)
	slices.Sort(sortedVals)
	diff := make([]int, 0)
	for i, v := range sortedVals {
		if v != originVals[i] {
			diff = append(diff, v)
		}
	}

	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val == diff[0] {
			node.Val = diff[1]
		} else if node.Val == diff[1] {
			node.Val = diff[0]
		}

		travel(node.Left)
		travel(node.Right)
	}

	travel(root)

}

func treeToSlice(root *TreeNode) []int {
	vals := make([]int, 0)

	var pushValues func(node *TreeNode)
	pushValues = func(node *TreeNode) {
		if node == nil {
			return
		}

		pushValues(node.Left)
		vals = append(vals, node.Val)
		pushValues(node.Right)
	}

	pushValues(root)

	return vals
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=end
