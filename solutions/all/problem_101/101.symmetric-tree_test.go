package problem101

import (
	"solutions/all/util"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		nodes []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nodes: []int{1, 2, 2, 3, 4, 4, 3},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nodes: []int{1, 2, 2, -1, 3, -1, 3},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				nodes: []int{2, 3, 3, 4, 5, -1, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := util.ConvertTree(tt.args.nodes)
			if got := isSymmetric(root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
