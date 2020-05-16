package str

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "abc", args: args{str: "abc"}, want: []byte{'a', 'b', 'c'}},
		{name: "abca", args: args{str: "abca"}, want: []byte{'a', 'b', 'c'}},
		{name: "abbca", args: args{str: "abca"}, want: []byte{'a', 'b', 'c'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
