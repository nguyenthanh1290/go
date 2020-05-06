package bitmnp

import "testing"

func Test_isSet(t *testing.T) {
	type args struct {
		n int
		i uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "20 = (10100)2",
			args: args{
				n: 20,
				i: 2,
			},
			want: true,
		},
		{
			name: "20 = (10100)2",
			args: args{
				n: 20,
				i: 1,
			},
			want: false,
		},
		{
			name: "1482 = (10111001010)2",
			args: args{
				n: 1482,
				i: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSet(tt.args.n, tt.args.i); got != tt.want {
				t.Errorf("isSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
