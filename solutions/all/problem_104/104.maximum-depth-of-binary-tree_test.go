package problem_104

import (
	"solutions/all/util"
	. "solutions/all/util"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		nodes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nodes: []int{3, 9, 20, -1, -1, 15, 7},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nodes: []int{1, -1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := util.ConvertTree(tt.args.nodes)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
