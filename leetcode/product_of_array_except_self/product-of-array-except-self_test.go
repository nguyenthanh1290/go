package array

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "",
			args: args{
				nums: []int{4, 5, 1, 8, 2},
			},
			want: []int{80, 64, 320, 40, 160},
		},
		{
			name: "",
			args: args{
				nums: []int{4, 5, 1, 8, 2, 10, 6},
			},
			want: []int{4800, 3840, 19200, 2400, 9600, 1920, 3200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
