package lru

import "container/list"

// LRU Cache
type Cache struct {
	maxBytes int64 // allow max memmory
	nbytes   int64 // current ocupy memory
	ll       *list.List
	cache    map[string]*list.Element
	OnEviced func(key string, value Value)
}

// node data structure
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEviced func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		OnEviced: onEviced,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	// search from map, if ok move to list's end
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// Remove oldest item
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEviced != nil {
			c.OnEviced(kv.key, kv.value)
		}
	}
}

// Add adds a value to cache
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		// update
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes -= int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}
