package bitmnp

import "testing"

func Test_countOnes(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 = (1)2",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "23 = (10111)2",
			args: args{
				x: 23,
			},
			want: 4,
		},
		{
			name: "30 = (11110)2",
			args: args{
				x: 30,
			},
			want: 4,
		},
		{
			name: "256 = (100000000)2",
			args: args{
				x: 256,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOnes(tt.args.x); got != tt.want {
				t.Errorf("countOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
