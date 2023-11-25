package cache

import "sync"

// SyncMap sync.Map with type parameters.
type SyncMap[T any] struct {
	syncMap sync.Map
}

func (sm *SyncMap[T]) Get(key string) (T, bool) {
	val, ok := sm.syncMap.Load(key)
	if !ok {
		return *new(T), false
	}

	return val.(T), true
}

func (sm *SyncMap[T]) Set(key string, value T) {
	sm.syncMap.Store(key, value)
}
