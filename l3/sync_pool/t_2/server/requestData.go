package server

import "sync"

type RequestData struct {
	data map[string]any
	mu   sync.Mutex
}

func (r *RequestData) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for key := range r.data {
		delete(r.data, key)
	}
}
