package data_structure

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestListInsert(t *testing.T) {
	type args struct {
		value interface{}
		index int
	}
	tests := []struct {
		name    string
		l       *List
		args    args
		want    *List
		wantErr bool
	}{
		{name: "_success_insert_at_middle",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{value: 11, index: 1},
			want: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 11,
						Next: &Item{
							Value: 10,
							Next: &Item{
								Value: 2,
								Next:  nil,
							},
						},
					},
				},
				Len: 4,
			},
			wantErr: false,
		},
		{name: "_success_insert_at_first",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 2,
						Next:  nil,
					},
				},
				Len: 2,
			},
			args: args{value: 10, index: 0},
			want: &List{
				Root: &Item{
					Value: 10,
					Next: &Item{
						Value: 1,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			wantErr: false,
		},
		{name: "_success_insert_at_last",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{value: 11, index: 3},
			want: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next: &Item{
								Value: 11,
								Next:  nil,
							},
						},
					},
				},
				Len: 4,
			},
			wantErr: false,
		},
		{name: "_error_out_of_length",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{value: 11, index: 4},
			want: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Insert(tt.args.value, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("List.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				bGot, _ := json.Marshal(got)
				bWant, _ := json.Marshal(tt.want)
				t.Errorf("List.Insert() = %v, want %v", string(bGot), string(bWant))
			}
		})
	}
}

func TestListRemove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		l       *List
		args    args
		want    *List
		wantErr bool
	}{
		{name: "_success_remove_at_middle",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{index: 1},
			want: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 2,
						Next:  nil,
					},
				},
				Len: 2,
			},
			wantErr: false,
		},
		{name: "_success_remove_at_first",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{index: 0},
			want: &List{
				Root: &Item{
					Value: 10,
					Next: &Item{
						Value: 2,
						Next:  nil,
					},
				},
				Len: 2,
			},
			wantErr: false,
		},
		{name: "_success_remove_at_last",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{index: 2},
			want: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next:  nil,
					},
				},
				Len: 2,
			},
			wantErr: false,
		},
		{name: "_error_out_of_length",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			args: args{index: 3},
			want: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 10,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
				},
				Len: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Remove(tt.args.index)

			if (err != nil) != tt.wantErr {
				t.Errorf("List.Remove(%d) error = %v, wantErr %v", tt.args.index, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				bGot, _ := json.Marshal(got)
				bWant, _ := json.Marshal(tt.want)
				t.Errorf("List.Remove(%d) = %v, want %v", tt.args.index, string(bGot), string(bWant))
			}
		})
	}
}

func TestListHas(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		l    *List
		args args
		want bool
	}{
		{name: "_nil_list",
			l:    &List{},
			args: args{value: 0},
			want: false,
		},
		{name: "_found_value",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 2,
						Next:  nil,
					},
				},
				Len: 2,
			},
			args: args{value: 2},
			want: true,
		},
		{name: "_not_found_value",
			l: &List{
				Root: &Item{
					Value: 1,
					Next: &Item{
						Value: 2,
						Next:  nil,
					},
				},
				Len: 2,
			},
			args: args{value: 5},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Has(tt.args.value); got != tt.want {
				t.Errorf("List.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
