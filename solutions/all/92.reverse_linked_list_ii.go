package problem_92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newHead := head

	if left > 1 {
		beforeLeft := head
		for i := 0; i < left-2; i++ {
			beforeLeft = beforeLeft.Next
		}

		beforeLeft.Next = reverseBetween(beforeLeft.Next, 1, right-left+1)
	} else {
		pointer := head
		for i := left; i < right; i++ {
			pointer = pointer.Next
		}
		end := pointer
		afterRight := end.Next
		end.Next = nil

		reversed := reverse(newHead)
		newHead = reversed
		pointer = reversed
		for pointer.Next != nil {
			pointer = pointer.Next
		}
		pointer.Next = afterRight
	}

	return newHead

}

func reverse(head *ListNode) *ListNode {
	newHead := head
	pointer := head.Next
	newHead.Next = nil
	for pointer != nil {
		next := pointer.Next
		curHead := newHead
		newHead = pointer
		newHead.Next = curHead
		pointer = next

	}

	return newHead
}
