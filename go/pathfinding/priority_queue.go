package main

import (
	"container/heap"
)

type PriorityQueue[T any] struct {
	items      []T
	comparator func(a, b T) bool
}

func NewPriorityQueue[T any](comparator func(a, b T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		items:      []T{},
		comparator: comparator,
	}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.comparator(pq.items[i], pq.items[j])
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue[T]) Push(item any) {
	pq.items = append(pq.items, item.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Top() T {
	if pq.Len() == 0 {
		var zeroValue T
		return zeroValue
	}
	return pq.items[0]
}

func (pq *PriorityQueue[T]) Add(item T) {
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) Remove() T {
	return heap.Pop(pq).(T)
}
