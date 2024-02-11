package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func Test_convertTree(t *testing.T) {
// 	nodes := []int{1, 3, -1, -1, 2}
// 	root := convertTree(nodes)
// 	vals := treeToSlice(root)
// 	fmt.Printf("vals: %v\n", vals)
// }

func ConvertTree(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	result := make([]*TreeNode, 0)
	for i, v := range nodes {
		if v == -1 {
			result = append(result, nil)
		} else {
			node := &TreeNode{Val: v}
			result = append(result, node)
			if i > 0 {
				if i%2 != 0 {
					result[(i-1)/2].Left = node
				} else {
					result[i/2-1].Right = node
				}
			}
		}
	}

	return result[0]
}

func TreeToSlice(root *TreeNode) []int {
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
