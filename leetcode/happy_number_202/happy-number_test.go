package happy_number_202

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy number",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "happy number",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "happy number",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "not a happy number",
			args: args{
				n: 20,
			},
			want: false,
		},
		{
			name: "not a happy number",
			args: args{
				n: 116,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
