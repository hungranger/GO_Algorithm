package searching

import "testing"

func TestLinearSearch(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "_found",
			args: args{
				arr: []int{3, 5, 6, 9, 7, 8, 5, 2},
				key: 9,
			},
			want: 3,
		},
		{name: "_not_found",
			args: args{
				arr: []int{3, 5, 6, 9, 7, 8, 5, 2},
				key: 27,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("LinearSearch(%d) = %v, want %v", tt.args.key, got, tt.want)
			}
		})
	}
}
