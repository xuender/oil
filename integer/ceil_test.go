package integer

import "testing"

func TestCeil(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1/2", args: args{dividend: 1, divisor: 2}, want: 1},
		{name: "4/2", args: args{dividend: 4, divisor: 2}, want: 2},
		{name: "5/2", args: args{dividend: 5, divisor: 2}, want: 3},
		{name: "10/3", args: args{dividend: 10, divisor: 3}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ceil(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}
