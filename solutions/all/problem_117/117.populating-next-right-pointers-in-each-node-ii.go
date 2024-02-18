package problem117

import . "solutions/all/util"

/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	links := []*Node{}
	var conn func(node *Node, depth int)
	conn = func(node *Node, depth int) {
		if node == nil {
			return
		}

		for len(links)-1 < depth {
			links = append(links, nil)
		}

		if node.Right != nil {
			node.Right.Next = links[depth]
			links[depth] = node.Right
		}

		if node.Left != nil {
			node.Left.Next = links[depth]
			links[depth] = node.Left
		}

		conn(node.Right, depth+1)
		conn(node.Left, depth+1)
	}
	conn(root, 0)

	return root
}

// @lc code=end
