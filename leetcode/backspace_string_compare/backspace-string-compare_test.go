package strcompare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				S: "ab#c",
				T: "ad#c",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				S: "ab##",
				T: "c#d#",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				S: "a##c",
				T: "#a#c",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				S: "a#c",
				T: "b",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				S: "ab",
				T: "b",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				S: "xywrrmp",
				T: "xywrrmu#p",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare(%v, %v) = %v, want %v", tt.args.S, tt.args.T, got, tt.want)
			}
		})
	}
}
