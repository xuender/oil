package integer

import "testing"

func TestDiv(t *testing.T) {
	type args struct {
		divisor int
		num     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "3/2", args: args{divisor: 3, num: 2}, want: 2},
		{name: "4/2", args: args{divisor: 4, num: 2}, want: 2},
		{name: "11/2", args: args{divisor: 11, num: 2}, want: 6},
		{name: "11/5", args: args{divisor: 11, num: 5}, want: 2},
		{name: "7/3", args: args{divisor: 7, num: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.divisor, tt.args.num); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
