package util

import (
	"reflect"
	"testing"
)

func TestConvertTree(t *testing.T) {
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
				nodes: []int{1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4},
			},
			want: []int{1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTree(tt.args.nodes); !reflect.DeepEqual(TreeToSlice(got), tt.want) {
				t.Errorf("ConvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
