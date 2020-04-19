package fibonacci

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
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
				n: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				n: 20,
			},
			want: 6765,
		},
		{
			name: "",
			args: args{
				n: 35,
			},
			want: 9227465,
		},
		{
			name: "",
			args: args{
				n: 100,
			},
			want: 3736710778780434371,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
