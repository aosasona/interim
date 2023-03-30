package interim

import (
	"fmt"
	"sync"
)

const (
	ERR_INVALID_KEY      = "invalid or blank key provided: %s"
	ERR_NO_VALUE         = "no value pair provided for key `%s`"
	ERR_KEY_ALREADY_SET  = "key `%s` is already set, use db.Update() to override current value"
	ERR_FAILED_TO_ENCODE = "failed to encode value for key `%s` to bytes"
)

type Interim struct {
	data  map[string][]byte
	cache *lruCache
	*sync.RWMutex
}

type Config struct {
	CacheSize int
}

func New(config Config) *Interim {
	cache := newLRUCache(config.CacheSize)

	return &Interim{
		data:  make(map[string][]byte),
		cache: cache,
	}
}

func (i *Interim) Set(key string, value any) error {
	i.Lock()
	defer i.Unlock()

	if key == "" {
		return fmt.Errorf(ERR_INVALID_KEY, key)
	}

	if value == nil {
		return fmt.Errorf(ERR_NO_VALUE, key)
	}

	if _, exists := i.data[key]; exists {
		return fmt.Errorf(ERR_KEY_ALREADY_SET, key)
	}

	encodedValue, err := encodeToByte(value)
	if err != nil {
		return fmt.Errorf(ERR_FAILED_TO_ENCODE, key)
	}

	i.data[key] = encodedValue
	i.cache.put(key, value)

	return nil
}
