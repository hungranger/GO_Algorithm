package sorting

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "_one_element_array",
			args: args{left: []int{5}, right: []int{2}},
			want: []int{2, 5},
		},
		{name: "_two_element_array",
			args: args{left: []int{2, 5}, right: []int{1, 7}},
			want: []int{1, 2, 5, 7},
		},
		{name: "_four_element_array",
			args: args{left: []int{1, 2, 5, 7}, right: []int{3, 4, 6, 8}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "_positive_array",
			args: args{arr: []int{1, 5, 6, 7, 8, 3, 4, 2}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{name: "_negative_array",
			args: args{arr: []int{1, 5, -6, -7, 0, 9, 10, 3, 4}},
			want: []int{-7, -6, 0, 1, 3, 4, 5, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
