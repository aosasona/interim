package interim

import (
	"fmt"
	"sync"
)

const (
	ERR_INVALID_KEY      = "invalid or blank key provided: %s"
	ERR_NO_VALUE         = "no value pair provided for key `%s`"
	ERR_KEY_NOT_FOUND    = "no record found for key `%s`"
	ERR_FAILED_TO_ENCODE = "failed to encode value for key `%s` to bytes"
	ERR_FAILED_TO_DECODE = "failed to decode value for key `%s`"
)

type interim struct {
	data  map[string][]byte
	cache *lruCache
	sync.RWMutex
}

type Config struct {
	CacheSize int
}

func New(config Config) *interim {
	cache := newLRUCache(config.CacheSize)
	return &interim{
		data:  make(map[string][]byte),
		cache: cache,
	}
}

func (i *interim) Get(key string, result interface{}) error {
	var (
		data    []byte
		ok, hit bool
	)

	i.RLock()
	defer i.RUnlock()

	if key == "" {
		return fmt.Errorf(ERR_INVALID_KEY, key)
	}

	raw, hit := i.cache.get(key)
	if hit {
		data, ok = raw.([]byte)
	} else {
		data, ok = i.data[key]
	}

	if !ok {
		return fmt.Errorf(ERR_KEY_NOT_FOUND, key)
	}

	err := decodeFromByte(data, result)

	if err != nil {
		return fmt.Errorf(ERR_FAILED_TO_DECODE, key)
	}

	return nil
}

func (i *interim) Set(key string, value any) error {
	i.Lock()
	defer i.Unlock()

	if key == "" {
		return fmt.Errorf(ERR_INVALID_KEY, key)
	}

	if value == nil {
		return fmt.Errorf(ERR_NO_VALUE, key)
	}

	encodedValue, err := encodeToByte(value)
	if err != nil {
		return fmt.Errorf(ERR_FAILED_TO_ENCODE, key)
	}

	i.data[key] = encodedValue
	i.cache.put(key, encodedValue)

	return nil
}

func (i *interim) Delete(key string) error {
	if key == "" {
		return fmt.Errorf(ERR_INVALID_KEY, key)
	}
	i.cache.remove(key)
	delete(i.data, key)
	return nil
}

func (i *interim) Exists(key string) bool {
	i.RLock()
	defer i.RUnlock()
	if key == "" {
		return false
	}
	_, ok := i.data[key]
	return ok
}

func (i *interim) Len() int {
	return len(i.data)
}
