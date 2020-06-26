// Package syncbychan https://blogtitle.github.io/go-advanced-concurrency-patterns-part-3-channels/
package syncbychan

type Item = interface{}

type Pool struct {
	buf   chan Item
	alloc func() Item
	clean func(Item) Item
}

func NewPool(size int, alloc func() Item, clean func(Item) Item) *Pool {
	return &Pool{
		buf:   make(chan Item, size),
		alloc: alloc,
		clean: clean,
	}
}

func (p *Pool) Get() Item {
	select {
	case i := <-p.buf:
		if p.clean != nil {
			return p.clean(i)
		}
		return i
	default:
		// Pool is empty, allocate new instance.
		return p.alloc()
	}
}

func (p *Pool) Put(x Item) {
	select {
	case p.buf <- x:
	default:
		// Pool is full, garbage-collect item.
	}
}
