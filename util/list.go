package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToSlice(head *ListNode) []int {
	vals := make([]int, 0)
	cur := head
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}

	return vals
}

func SliceToList(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}

	head := &ListNode{Val: nodes[0]}
	cur := head
	for i := 1; i < len(nodes); i++ {
		link := &ListNode{Val: nodes[i]}
		cur.Next = link
		cur = link
	}

	return head
}
