package bitwise

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				m: 5,
				n: 7,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				m: 0,
				n: 1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				m: 1,
				n: 11,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
