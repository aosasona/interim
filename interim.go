package interim

import "sync"

type Interim struct {
	data map[string][]byte
	pool sync.Pool
	*sync.RWMutex
}

func New() *Interim {
	return &Interim{
		data: make(map[string][]byte),
	}
}
