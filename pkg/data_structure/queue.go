package data_structure

import "errors"

type Queue struct {
	List
}

func NewQueue() *Queue {
	s := &Queue{}
	s.Len = 0
	return s
}

func (s *Queue) Enqueue(value interface{}) (*Queue, error) {
	//insert_at_last
	s.Insert(value, s.Len)
	return s, nil
}

func (s *Queue) Dequeue() (*Queue, error) {
	if s.Len == 0 {
		return s, errors.New("Empty Queue")
	} else {
		//remove_at_first
		s.Remove(0)
	}

	return s, nil
}
