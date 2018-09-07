package data_structure

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStackPush(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		s       *Stack
		args    args
		want    *Stack
		wantErr bool
	}{
		{name: "_success_existed_stack",
			s: &Stack{
				List: List{
					Root: &Item{
						Value: 1,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
					Len: 2,
				},
			},
			args: args{value: 10},
			want: &Stack{
				List: List{
					Root: &Item{
						Value: 1,
						Next: &Item{
							Value: 2,
							Next: &Item{
								Value: 10,
								Next:  nil,
							},
						},
					},
					Len: 3,
				},
			},
			wantErr: false,
		},
		{name: "_success_existed_stack",
			s:    &Stack{List: List{Len: 0}},
			args: args{value: 10},
			want: &Stack{
				List: List{
					Root: &Item{
						Value: 10,
						Next:  nil,
					},
					Len: 1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Push(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Push(%d) error = %v, wantErr %v", tt.args.value, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				bGot, _ := json.Marshal(got)
				bWant, _ := json.Marshal(tt.want)
				t.Errorf("Stack.Push(%d) = %v, want %v", tt.args.value, string(bGot), string(bWant))
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stack
		want    *Stack
		wantErr bool
	}{
		{name: "_error_pop_empty_stack",
			s:       &Stack{List: List{Len: 0}},
			want:    &Stack{List: List{Len: 0}},
			wantErr: true,
		},
		{name: "_success_pop_existed_stack",
			s: &Stack{
				List: List{
					Root: &Item{
						Value: 1,
						Next: &Item{
							Value: 2,
							Next:  nil,
						},
					},
					Len: 2,
				},
			},
			want: &Stack{
				List: List{
					Root: &Item{
						Value: 1,
						Next:  nil,
					},
					Len: 1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				bGot, _ := json.Marshal(got)
				bWant, _ := json.Marshal(tt.want)
				t.Errorf("Stack.Pop() = %v, want %v", string(bGot), string(bWant))
			}
		})
	}
}
