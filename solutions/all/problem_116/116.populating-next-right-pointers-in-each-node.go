package problem116

import (
	. "solutions/all/util"
)

/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
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
	head := root
	for head.Right != nil {
		links = append(links, nil)
		head = head.Right
	}
	conn(root, 0, links)

	return root
}

func conn(node *Node, depth int, links []*Node) {
	if node == nil {
		return
	}

	if node.Right != nil {
		node.Right.Next = links[depth]
		node.Left.Next = node.Right
		links[depth] = node.Left
		conn(node.Right, depth+1, links)
		conn(node.Left, depth+1, links)
	}
}

// @lc code=end
