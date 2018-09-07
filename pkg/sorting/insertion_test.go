package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "_positive_array",
			args: args{arr: []int{1, 5, 6, 7, 8, 9, 10, 3, 4, 2, 15}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15},
		},
		{name: "_negative_array",
			args: args{arr: []int{1, 5, -6, -7, 0, 9, 10, 3, 4}},
			want: []int{-7, -6, 0, 1, 3, 4, 5, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
