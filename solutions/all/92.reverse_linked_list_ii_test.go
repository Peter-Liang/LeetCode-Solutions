package problem_92

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseFromMiddle(t *testing.T) {
	head := createList([]int{1, 2, 3, 4, 5})
	head = reverseBetween(head, 2, 4)

	assert.Equal(t, []int{1, 4, 3, 2, 5}, toList(head))
}

func TestReverseTillEnd(t *testing.T) {
	head := createList([]int{1, 2, 3, 4, 5})
	head = reverseBetween(head, 2, 5)

	assert.Equal(t, []int{1, 5, 4, 3, 2}, toList(head))
}

func TestReverseFromStart(t *testing.T) {
	head := createList([]int{1, 2, 3, 4, 5})
	head = reverseBetween(head, 1, 3)

	assert.Equal(t, []int{3, 2, 1, 4, 5}, toList(head))
}

func TestReverseAll(t *testing.T) {
	head := createList([]int{1, 2, 3, 4, 5})
	head = reverseBetween(head, 1, 5)

	assert.Equal(t, []int{5, 4, 3, 2, 1}, toList(head))
}

func toList(head *ListNode) []int {
	result := make([]int, 0)
	pointer := head
	for pointer != nil {
		result = append(result, pointer.Val)
		pointer = pointer.Next
	}

	return result
}

func createList(list []int) *ListNode {
	head := ListNode{}
	pointer := &head
	for _, val := range list {
		pointer.Next = &ListNode{Val: val}
		pointer = pointer.Next
	}

	return head.Next
}
