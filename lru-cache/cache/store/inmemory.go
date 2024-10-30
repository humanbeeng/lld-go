package store

import (
	"sync"
)

type InMemoryStore[K, V comparable] struct {
	store    map[K]V
	capacity int
}

var instance *InMemoryStore[string, string]
var once sync.Once

func GetInMemoryStore(capacity int) *InMemoryStore[string, string] {
	once.Do(func() {
		inmemStore := newInMemoryStore[string, string](capacity)
		instance = &inmemStore
	})
	return instance
}

func newInMemoryStore[K, V comparable](capacity int) InMemoryStore[K, V] {
	return InMemoryStore[K, V]{
		store:    make(map[K]V),
		capacity: capacity,
	}
}

func (s *InMemoryStore[K, V]) Get(key K) (V, error) {
	val, ok := s.store[key]

	if !ok {
		return val, ErrKeyNotFound
	}

	return val, nil
}

func (s *InMemoryStore[K, V]) Set(key K, value V) error {
	if s.capacity == len(s.store) {
		return ErrCapacityExceeded
	}

	s.store[key] = value
	return nil
}

func (s *InMemoryStore[K, V]) Del(key K) error {
	_, ok := s.store[key]
	if !ok {
		return ErrKeyNotFound
	}

	delete(s.store, key)
	return nil
}
