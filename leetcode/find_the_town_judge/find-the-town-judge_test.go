package judge

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		N     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				N:     2,
				trust: [][]int{{1, 2}},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				N: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
				},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				N: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
					{3, 1},
				},
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				N: 3,
				trust: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				N: 4,
				trust: [][]int{
					{1, 3},
					{1, 4},
					{2, 3},
					{2, 4},
					{4, 3},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.N, tt.args.trust); got != tt.want {
				t.Errorf("findJudge(%v) = %v, want %v", tt.args.trust, got, tt.want)
			}
		})
	}
}
