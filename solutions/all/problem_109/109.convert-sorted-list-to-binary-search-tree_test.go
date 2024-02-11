package problem109

import (
	"reflect"
	. "solutions/all/util"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
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
				nodes: []int{-10, -3, 0, 5, 9},
			},
			want: []int{0, -3, 9, -10, -1, 5},
		},
		{
			name: "case 2",
			args: args{
				nodes: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(SliceToList(tt.args.nodes)); !reflect.DeepEqual(got, ConvertTree(tt.want)) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
