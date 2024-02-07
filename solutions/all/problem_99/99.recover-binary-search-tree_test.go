package problem99

import (
	"reflect"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		nodes []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nodes: []int{1, 3, -1, -1, 2},
			},
			want: []int{3, 1, -1, -1, 2},
		},
		{
			name: "case 2",
			args: args{
				nodes: []int{3, 1, 4, -1, -1, 2},
			},
			want: []int{2, 1, 4, -1, -1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := convertTree(tt.args.nodes)
			recoverTree(root)
			// got := treeToSlice(root)
			want := convertTree(tt.want)

			if !reflect.DeepEqual(root, want) {
				t.Errorf("not equal %v, want %v", root, want)
			}
		})
	}
}

// func Test_convertTree(t *testing.T) {
// 	nodes := []int{1, 3, -1, -1, 2}
// 	root := convertTree(nodes)
// 	vals := treeToSlice(root)
// 	fmt.Printf("vals: %v\n", vals)
// }

func convertTree(nodes []int) *TreeNode {
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
