package problem_97

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{"aabcc", "dbbca", "aadbbcbcac"},
			want: true,
		},
		{
			name: "case 2",
			args: args{"aabcc", "dbbca", "aadbbbaccc"},
			want: false,
		},
		{
			name: "case 3",
			args: args{"", "", ""},
			want: true,
		},
		{
			name: "case 4",
			args: args{"a", "b", "a"},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				"abababababababababababababababababababababababababababababababababababababababababababababababababbb",
				"babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
				"abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
