package interim

import "container/list"

type lruCache struct {
	capacity int
	cache    map[any]*list.Element
	list     *list.List
}

type cacheEntry struct {
	key   string
	value any
}

func newLRUCache(capacity int) *lruCache {
	if capacity == 0 {
		// This is used to programmitically limit the size (we don't want to use make(..., capacity) because it will retain memory space we may not even use)
		capacity = 8
	}

	return &lruCache{
		capacity: capacity,
		cache:    make(map[any]*list.Element),
		list:     list.New(),
	}
}

func (l *lruCache) get(key string) (any, bool) {
	entry, hit := l.cache[key]
	if !hit {
		return nil, false
	}

	// Move the most recently accessed item to the front of the list IN CASE it is requested for again soon after that
	l.list.MoveToFront(entry)

	// Cast to cacheEntry type and return only the value (by default, list.Element contains both the key and the value)
	return entry.Value.(*cacheEntry).value, true
}

func (l *lruCache) put(key string, value any) {
	// If it exists (cache hit), update and move it to the front as the most recently accessed entry
	if entry, hit := l.cache[key]; hit {
		l.list.MoveToFront(entry)
		entry.Value.(*cacheEntry).value = value
		return
	}

	// If the list is full, remove the last element (which will also be our least recently accessed entry) to make space for the new one at the front
	if l.len() == l.capacity {
		entry := l.list.Back()
		delete(l.cache, entry.Value.(*cacheEntry).key)
		l.list.Remove(entry)
	}

	entry := l.list.PushFront(&cacheEntry{key, value})
	l.cache[key] = entry
	return
}

func (l *lruCache) len() int {
	return l.list.Len()
}
