package problem101

import (
	. "solutions/all/util"
)

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

func isSymmetric(root *TreeNode) bool {
	leftCount, rightCount := 1, 1
	leftChan := make(chan *TreeNode)
	rightChan := make(chan *TreeNode)

	var travel func(*TreeNode, bool)
	travel = func(node *TreeNode, isLeft bool) {
		if isLeft {
			leftChan <- node
		} else {
			rightChan <- node
		}
		if node != nil {
			if isLeft {
				leftCount++
				travel(node.Left, true)
				leftCount++
				travel(node.Right, true)
			} else {
				rightCount++
				travel(node.Right, false)
				rightCount++
				travel(node.Left, false)
			}
		}

		if isLeft {
			leftCount--
			if leftCount == 0 {
				close(leftChan)
			}
		} else {
			rightCount--
			if rightCount == 0 {
				close(rightChan)
			}
		}
	}

	go travel(root.Left, true)
	go travel(root.Right, false)

	for {
		left, leftOk := <-leftChan
		right, rightOk := <-rightChan

		if leftOk != rightOk {
			return false
		}

		if leftOk {
			if left == nil && right == nil {
				continue
			}
			if left == nil && right != nil {
				return false
			}

			if right == nil && left != nil {
				return false
			}

			if left.Val != right.Val {
				return false
			}
		} else {
			return true
		}
	}
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// @lc code=end
