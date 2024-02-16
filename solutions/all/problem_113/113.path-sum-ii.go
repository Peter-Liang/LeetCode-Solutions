package problem113

import . "solutions/all/util"

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	val := root.Val
	if root.Left == nil && root.Right == nil && targetSum == val {
		return [][]int{{root.Val}}
	}

	res := make([][]int, 0)
	res = append(res, pathSum(root.Left, targetSum-val)...)
	res = append(res, pathSum(root.Right, targetSum-val)...)
	for i, v := range res {
		res[i] = append([]int{val}, v...)
	}

	return res
}

// @lc code=end
