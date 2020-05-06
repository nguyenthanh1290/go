package bitmnp

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 64,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 11,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 256,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 1048576,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
