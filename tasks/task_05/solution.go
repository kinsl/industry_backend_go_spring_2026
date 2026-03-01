package main

type Cache[K comparable, V any] struct {
	items map[K]V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	var items map[K]V
	if capacity > 0 {
		items = make(map[K]V, capacity)
	}
	return &Cache[K, V]{
		items: items,
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	if c.items != nil {
		c.items[key] = value
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	val, ok := c.items[key]
	return val, ok
}
