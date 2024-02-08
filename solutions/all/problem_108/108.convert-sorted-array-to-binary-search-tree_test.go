package problem108

import (
	"reflect"
	"testing"

	. "solutions/all/util"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: []int{0, -3, 9, -10, -1, 5},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 3},
			},
			want: []int{3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := ConvertTree(tt.want)
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
