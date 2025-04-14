package cache

import "container/list"

type Cache[K comparable, V any] struct {
	total int64
	cnt   int64
	ll    *list.List
	cache map[K]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key K, value V)
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](total int64, onEvicted func(K, V)) *Cache[K, V] {
	return &Cache[K, V]{
		total:     total,
		ll:        list.New(),
		cache:     make(map[K]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache[K, V]) Get(key K) (V, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry[K, V])
		return kv.value, true
	}
	var zero V
	return zero, false
}

// RemoveOldest removes the oldest item
func (c *Cache[K, V]) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry[K, V])
		delete(c.cache, kv.key)
		c.cnt -= int64(1)
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (c *Cache[K, V]) Add(key K, value V) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry[K, V])
		c.cnt += int64(1)
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry[K, V]{key, value})
		c.cache[key] = ele
		c.cnt += int64(1)
	}
	if c.total != 0 && c.cnt > c.total {
		c.RemoveOldest()
	}
}

// Len returns the number of items in the cache.
func (c *Cache[K, V]) Len() int {
	return c.ll.Len()
}
