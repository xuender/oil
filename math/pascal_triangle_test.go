package math

import (
	"fmt"
	"reflect"
	"testing"
)

func Benchmark_PascalTriangle(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		PascalTriangle(1000)
	}
}

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
			name: "-1",
			args: args{layer: -1},
			want: []int{},
		},
		{
			name: "0",
			args: args{layer: 0},
			want: []int{1},
		},
		{
			name: "1",
			args: args{layer: 1},
			want: []int{1, 1},
		},
		{
			name: "2",
			args: args{layer: 2},
			want: []int{1, 2, 1},
		},
		{
			name: "3",
			args: args{layer: 3},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "4",
			args: args{layer: 4},
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

func ExamplePascalTriangle() {
	fmt.Println(PascalTriangle(8))

	// Output:
	// [1 8 28 56 70 56 28 8 1]
}

func ExamplePascalSum() {
	/*
				赌徒费马和帕斯卡，每一盘里，他俩的赢的机会相等。
				他俩各拿出相同金额的钱作为赌注，约定谁先赢到某个（假设是10）盘数，赌注就全部归谁。
				不料，这时发生了某事，他们必须结束赌局并离开。
				此时，两个人谁也没赢到10盘，那么这个赌注的钱应该怎么分呢？
				当然，此时赢得多的人应该相应地拿的赌注多。
				可是，多少才算是公平呢？

				当赌局进行到 费马8比帕斯卡7
		    还需要最多4盘才能决出胜负
		    费马需要胜利2次，帕斯卡需要胜利3次
	*/
	sum := PascalSum(4, 4)
	fmt.Printf("费马=%d/%d, 帕斯卡=%d/%d\n", PascalSum(4, 2), sum, PascalSum(4, 3), sum)

	// Output:
	// 费马=11/16, 帕斯卡=5/16
}
