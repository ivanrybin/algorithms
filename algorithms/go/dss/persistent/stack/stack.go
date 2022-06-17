package stack

import (
	"fmt"
)

type Stack interface {
	Push(v int)
	Pop() (int, error)
	Empty() bool

	Version() int
	List(version int) ([]int, error)
}

func New() *stack {
	return &stack{heads: nil}
}

type item struct {
	value int
	next  *item
}

func (i *item) list() []int {
	curr := i
	var list []int
	for curr != nil {
		list = append(list, curr.value)
		curr = curr.next
	}
	return list
}

type stack struct {
	heads []*item
}

func (s *stack) Push(v int) {
	var curr *item
	if !s.empty() {
		curr = s.heads[len(s.heads)-1]
	}
	s.heads = append(s.heads, &item{
		value: v,
		next:  curr,
	})
}

func (s *stack) Pop() (int, error) {
	if s.empty() {
		return 0, fmt.Errorf("empty")
	}
	curr := s.heads[len(s.heads)-1]
	s.heads = append(s.heads, curr.next)
	return curr.value, nil
}

func (s *stack) Empty() bool {
	return s.empty()
}

func (s *stack) empty() bool {
	if len(s.heads) == 0 {
		return true
	}
	return s.heads[len(s.heads)-1] == nil
}

func (s *stack) Version() int {
	return len(s.heads) - 1
}

func (s *stack) List(version int) ([]int, error) {
	if version < 0 || version >= len(s.heads) {
		return nil, fmt.Errorf("invalid version")
	}
	return s.heads[version].list(), nil
}
