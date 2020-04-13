package array

import "testing"

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 0},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0, 0, 1},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0, 0, 1, 1},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
