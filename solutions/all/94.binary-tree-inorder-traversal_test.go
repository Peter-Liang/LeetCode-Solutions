package problem_94

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	result := inorderTraversal(&root)

	assert.Equal(t, []int{1, 3, 2}, result)
}
