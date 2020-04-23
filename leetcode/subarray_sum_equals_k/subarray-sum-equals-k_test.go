package subarray

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1},
				k:    3,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1},
				k:    1,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 3, 4, 7, 5},
				k:    11,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 2, 1},
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum(%+v, %v) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}
