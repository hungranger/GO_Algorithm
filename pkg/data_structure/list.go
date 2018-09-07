package data_structure

import "errors"

type List struct {
	Root *Item
	Len  int
}

type Item struct {
	Value interface{}
	Next  *Item
	// list  *List
}

func NewList() *List {
	l := &List{}
	l.Len = 0
	return l
}

func (l *List) Remove(index int) (*List, error) {
	if index > l.Len-1 || index < 0 {
		return l, errors.New("Out of length")
	}

	preIndexItem := l.Root

	if index == 0 {
		l.Root = preIndexItem.Next
		preIndexItem = nil
	} else {
		for i := 0; i < index-1; i++ {
			preIndexItem = preIndexItem.Next
		}
		indexItem := preIndexItem.Next
		nextIndexItem := indexItem.Next

		preIndexItem.Next = nextIndexItem
		indexItem = nil
	}

	l.Len--

	return l, nil
}

func (l *List) Insert(value interface{}, index int) (*List, error) {
	if index > l.Len || index < 0 {
		return l, errors.New("Out of length")
	}

	newItem := &Item{value, nil}
	if index == 0 {
		newItem.Next = l.Root
		l.Root = newItem
	} else {
		indexItem := l.Root
		for i := 0; i < index-1; i++ {
			indexItem = indexItem.Next
		}

		newItem.Next = indexItem.Next
		indexItem.Next = newItem
	}

	l.Len++

	return l, nil
}

func (l *List) Has(value interface{}) bool {
	root := l.Root

	if root == nil {
		return false
	}

	for {
		if root.Value == value {
			return true
		} else {
			if root.Next != nil {
				root = root.Next
			} else {
				return false
			}
		}
	}
}

func (l *List) Length() int {
	return l.Len
}
