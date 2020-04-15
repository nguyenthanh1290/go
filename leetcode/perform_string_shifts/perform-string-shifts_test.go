package string

import "testing"

func Test_stringShift(t *testing.T) {
	type args struct {
		s     string
		shift [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s:     "abc",
				shift: [][]int{{0, 1}, {1, 2}},
			},
			want: "cab",
		},
		{
			name: "",
			args: args{
				s:     "abcdefg",
				shift: [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}},
			},
			want: "efgabcd",
		},
		{
			name: "",
			args: args{
				s:     "abc",
				shift: [][]int{{0, 1}, {0, 1}, {0, 1}},
			},
			want: "abc",
		},
		{
			name: "",
			args: args{
				s:     "abc",
				shift: [][]int{{1, 1}, {1, 1}, {1, 1}},
			},
			want: "abc",
		},
		{
			name: "",
			args: args{
				s:     "mecsk",
				shift: [][]int{{1, 4}, {0, 5}, {0, 4}, {1, 1}, {1, 5}},
			},
			want: "kmecs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringShift(tt.args.s, tt.args.shift); got != tt.want {
				t.Errorf("stringShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
