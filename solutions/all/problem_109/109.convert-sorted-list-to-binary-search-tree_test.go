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
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(SliceToList(tt.args.nodes)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
