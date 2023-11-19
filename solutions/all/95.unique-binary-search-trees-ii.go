package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	s := fmt.Sprintf("[%d", n.Val)

	if n.Left != nil {
		s += fmt.Sprintf(", left: %s", n.Left)
	}

	if n.Right != nil {
		s += fmt.Sprintf(", right: %s", n.Right)
	}

	s += "]"

	return s
}

func main() {
	a := []int{4}
	for _, v := range a {
		res := generateTrees(v)
		for i, tn := range res {
			fmt.Printf("%d -> %v\n", i, tn)
		}
	}
}

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {
	orig := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		orig = append(orig, i)
	}

	res := create(orig)
	for i := len(res) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if equals(res[i], res[j]) {
				res = append(res[:i], res[i+1:]...)
				break
			}
		}
	}

	return res
}

func equals(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil || a.Val != b.Val {
		return false
	}

	return equals(a.Left, b.Left) && equals(a.Right, b.Right)
}

func create(vals []int) []*TreeNode {
	if len(vals) == 0 {
		return []*TreeNode{}
	}

	if len(vals) == 1 {
		node := &TreeNode{Val: vals[0]}
		return []*TreeNode{node}
	}

	var ans []*TreeNode
	lists := getLists(vals)
	for _, list := range lists {
		root := &TreeNode{Val: list[0]}
		for _, v := range list[1:] {
			insert(root, v)
		}
		ans = append(ans, root)
	}

	return ans
}

func getLists(vals []int) [][]int {
	if len(vals) == 0 {
		return [][]int{}
	}

	if len(vals) == 1 {
		return [][]int{vals}
	}

	var ans [][]int
	for i, v := range vals {
		var remains []int
		remains = append(remains, vals[:i]...)
		remains = append(remains, vals[i+1:]...)

		res := getLists(remains)
		for _, list := range res {
			root := []int{v}
			ans = append(ans, append(root, list...))
		}
	}

	return ans
}

func insert(tree *TreeNode, node int) *TreeNode {
	if tree == nil {
		return &TreeNode{Val: node}
	}
	if tree.Val > node {
		tree.Left = insert(tree.Left, node)
	} else {
		tree.Right = insert(tree.Right, node)
	}

	return tree
}

// @lc code=end
