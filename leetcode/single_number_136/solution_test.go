package single_number_136

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				nums: []int{1, 1, 3, 2, 2},
			},
			want: 3,
		},
		{
			name: "normal: include 2 digits",
			args: args{
				nums: []int{1, 1, 3, 22, 22},
			},
			want: 3,
		},
		{
			name: "found no single number",
			args: args{
				nums: []int{1, 1, 3, 2, 2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberUsingXOR(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberUsingXOR() = %v, want %v", got, tt.want)
			}
			if got := singleNumberUsingMap(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberUsingMap() = %v, want %v", got, tt.want)
			}
			if got := singleNumberUsingMath(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberUsingMath() = %v, want %v", got, tt.want)
			}
		})
	}
}
