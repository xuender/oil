package integer

import "testing"

func TestPow(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1**1", args: args{1, 1}, want: 1},
		{name: "2**2", args: args{2, 2}, want: 4},
		{name: "2**3", args: args{2, 3}, want: 8},
		{name: "3**2", args: args{3, 2}, want: 9},
		{name: "3**3", args: args{3, 3}, want: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
