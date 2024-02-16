package problem115

import (
	"testing"
)

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "rabbbit",
				t: "rabit",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				s: "babgbag",
				t: "bag",
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				s: "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc",
				t: "bcddceeeebecbc",
			},
			want: 700531452,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
