package circular_buffer

import (
	"errors"
	"sync"
)

type T interface{}

type Buffer interface {
	IsEmpty() bool
	IsFull() bool
	Add(task interface{}) error
	Delete() (interface{}, error)
}

var (
	errFull   = errors.New("full")
	errNoTask = errors.New("no task")
)

type CircularBuffer struct {
	sync.Mutex
	taskQueue []T
	capacity  int
	head      int
	tail      int
	full      bool
}

func NewCircularBuffer(size int) *CircularBuffer {
	w := &CircularBuffer{
		taskQueue: make([]T, size),
		capacity:  size,
	}

	return w
}

func (s *CircularBuffer) IsEmpty() bool {
	return s.head == s.tail && !s.full
}

func (s *CircularBuffer) IsFull() bool {
	return s.full
}

func (s *CircularBuffer) Add(task interface{}) error {
	if s.IsFull() {
		return errFull
	}

	s.Lock()
	s.taskQueue[s.tail] = task.(T)
	s.tail = (s.tail + 1) % s.capacity
	s.full = s.head == s.tail
	s.Unlock()

	return nil
}

func (s *CircularBuffer) Delete() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errNoTask
	}

	s.Lock()
	data := s.taskQueue[s.head]
	s.full = false
	s.head = (s.head + 1) % s.capacity
	s.Unlock()

	return data, nil
}
