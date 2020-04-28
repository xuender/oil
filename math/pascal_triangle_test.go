package math

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	type args struct {
		layer int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0",
			args: args{layer: 0},
			want: []int{},
		},
		{
			name: "1",
			args: args{layer: 1},
			want: []int{1},
		},
		{
			name: "2",
			args: args{layer: 2},
			want: []int{1, 1},
		},
		{
			name: "3",
			args: args{layer: 3},
			want: []int{1, 2, 1},
		},
		{
			name: "4",
			args: args{layer: 4},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "5",
			args: args{layer: 5},
			want: []int{1, 4, 6, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalTriangle(tt.args.layer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PascalTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
