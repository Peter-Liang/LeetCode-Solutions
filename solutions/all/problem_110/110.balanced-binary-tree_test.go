package main

import (
	. "solutions/all/util"
	"testing"
)

func Test_isBalanced(t *testing.T) {
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
				nodes: []int{1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4},
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				nodes: []int{3, 9, 20, -1, -1, 15, 7},
			},
			want: true,
		},
		// {
		// 	name: "case 3",
		// 	args: args{
		// 		nodes: []int{1, 2, 2, 3, 3, -1, -1, 4, 4},
		// 	},
		// 	want: false,
		// },
		{
			name: "case 4",
			args: args{
				nodes: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(ConvertTree(tt.args.nodes)); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
