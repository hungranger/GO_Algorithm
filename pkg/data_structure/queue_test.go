package data_structure

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		s       *Queue
		args    args
		want    *Queue
		wantErr bool
	}{
		{name: "_success_existed_Queue",
			s: &Queue{
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
			want: &Queue{
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
		{name: "_success_existed_Queue",
			s:    &Queue{List: List{Len: 0}},
			args: args{value: 10},
			want: &Queue{
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
			got, err := tt.s.Enqueue(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				bGot, _ := json.Marshal(got)
				bWant, _ := json.Marshal(tt.want)
				t.Errorf("Queue.Enqueue(%d) = %v, want %v", tt.args.value, string(bGot), string(bWant))
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		s       *Queue
		want    *Queue
		wantErr bool
	}{
		{name: "_error_pop_empty_Queue",
			s:       &Queue{List: List{Len: 0}},
			want:    &Queue{List: List{Len: 0}},
			wantErr: true,
		},
		{name: "_success_pop_existed_Queue",
			s: &Queue{
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
			want: &Queue{
				List: List{
					Root: &Item{
						Value: 2,
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
			got, err := tt.s.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				bGot, _ := json.Marshal(got)
				bWant, _ := json.Marshal(tt.want)
				t.Errorf("Queue.Dequeue() = %v, want %v", string(bGot), string(bWant))
			}
		})
	}
}
