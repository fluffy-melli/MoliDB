package cache

import (
	"sync"
)

type Safe struct {
	mu    sync.RWMutex
	store map[string]any
}

func NewSafeMap() *Safe {
	return &Safe{
		store: make(map[string]any, 0),
	}
}

func (sm *Safe) Get(key string) (any, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.store[key]
	return value, exists
}

func (sm *Safe) Set(key string, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.store[key] = value
}

func (sm *Safe) Del(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.store, key)
}

func (s *Safe) GetKeys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.store))
	for key := range s.store {
		keys = append(keys, key)
	}
	return keys
}

func (s *Safe) GetStore() map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.store
}
