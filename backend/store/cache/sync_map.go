package cache

import "sync"

type SyncMap[T any] struct {
	m sync.Map
}

func (s *SyncMap[T]) Get(key string) (T, bool) {
	val, ok := s.m.Load(key)
	if !ok {
		return *new(T), false
	}

	return val.(T), true
}

func (s *SyncMap[T]) Set(key string, value T) {
	s.m.Store(key, value)
}
