package utils

import (
	"sync"
	"sync/atomic"
)

type CopyOnWriteSet[T comparable] struct {
	val atomic.Value // stores map[T]struct{}
	mu  sync.Mutex
}

func NewSet[T comparable]() *CopyOnWriteSet[T] {
	s := &CopyOnWriteSet[T]{}
	s.val.Store(make(map[T]struct{}))
	return s
}

func (s *CopyOnWriteSet[T]) Add(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	old := s.val.Load().(map[T]struct{})
	c := make(map[T]struct{}, len(old)+1)
	for k := range old {
		c[k] = struct{}{}
	}
	c[v] = struct{}{}
	s.val.Store(c)
}

func (s *CopyOnWriteSet[T]) Remove(v T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	old := s.val.Load().(map[T]struct{})
	if _, exists := old[v]; !exists {
		return
	}

	c := make(map[T]struct{}, len(old)-1)
	for k := range old {
		if k != v {
			c[k] = struct{}{}
		}
	}
	s.val.Store(c)
}

func (s *CopyOnWriteSet[T]) Contains(v T) bool {
	current := s.val.Load().(map[T]struct{})
	_, ok := current[v]
	return ok
}

func (s *CopyOnWriteSet[T]) Snapshot() []T {
	current := s.val.Load().(map[T]struct{})
	result := make([]T, 0, len(current))
	for k := range current {
		result = append(result, k)
	}
	return result
}

func (s *CopyOnWriteSet[T]) Size() int {
	current := s.val.Load().(map[T]struct{})
	return len(current)
}
