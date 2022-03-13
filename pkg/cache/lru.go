package cache

import (
	"container/ring"
)

type keyValue[K comparable, V any] struct {
	key   K
	value V
}

// LRU implements least recently used cache.
// Key equality is determined in the same way as in a map.
type LRU[K comparable, V any] struct {
	size  int
	cap   int
	head  *ring.Ring
	byKey map[K]*ring.Ring
}

// NewLRU creates empty LRU cache.
// cap must be positive.
func NewLRU[K comparable, V any](cap int) *LRU[K, V] {
	return &LRU[K, V]{
		cap:   cap,
		byKey: make(map[K]*ring.Ring),
	}
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	r, ok := c.byKey[key]
	if !ok {
		var v V
		return v, false
	}

	r.Prev().Link(r.Next())
	r.Link(c.head)
	c.head = r

	kv := r.Value.(*keyValue[K, V])
	return kv.value, true
}

func (c *LRU[K, V]) Set(key K, val V) {
	kv := &keyValue[K, V]{key, val}

	r, ok := c.byKey[key]
	if ok {
		r.Value = kv
		return
	}

	if c.size < c.cap {
		c.size++

		r := ring.New(1)
		c.byKey[key] = r
		r.Value = kv

		if c.head == nil {
			c.head = r
			return
		}

		r.Link(c.head)
		c.head = r
		return
	}

	tail := c.head.Prev()

	oldKV := tail.Value.(*keyValue[K, V])
	delete(c.byKey, oldKV.key)

	c.byKey[key] = tail
	tail.Value = kv

	c.head = tail
}

func (c *LRU[K, V]) Remove(key K) {
	r, ok := c.byKey[key]
	if !ok {
		return
	}

	c.size--
	delete(c.byKey, key)

	if c.size == 0 {
		c.head = nil
		return
	}

	if r == c.head {
		c.head = r.Next()
	}
	r.Prev().Link(r.Next())
}
