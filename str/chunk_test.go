package str

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		str  string
		size int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "123 | 2", args: args{str: "123", size: 1}, want: []string{"1", "2", "3"}},
		{name: "12 | 2", args: args{str: "12", size: 2}, want: []string{"12"}},
		{name: "456 | 2", args: args{str: "456", size: 2}, want: []string{"45", "6"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.str, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
