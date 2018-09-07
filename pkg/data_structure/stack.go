package data_structure

import "errors"

type Stack struct {
	List
}

func NewStack() *Stack {
	s := &Stack{}
	s.Len = 0
	return s
}

func (s *Stack) Push(value interface{}) (*Stack, error) {
	//insert_at_last
	s.Insert(value, s.Len)
	return s, nil
}

func (s *Stack) Pop() (*Stack, error) {
	if s.Len == 0 {
		return s, errors.New("Empty stack")
	} else {
		//remove_at_last
		s.Remove(s.Len - 1)
	}

	return s, nil
}
