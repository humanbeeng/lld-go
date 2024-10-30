package store

import (
	"context"
	"sync"
)

type Store interface {
	Get(key string) (string, error)
	Set(key string, val string) error
	Del(key string) error

	BeginTransaction() *Transaction
	KeyVersion(key string) int
}

type InMemoryStore struct {
	store   map[string]VersionedValue
	version int
	sync.RWMutex
}

var inMemStoreInstance *InMemoryStore
var once sync.Once

func GetInMemoryStore() *InMemoryStore {
	once.Do(func() {
		inMemStoreInstance = newInMemoryStore()
	})

	return inMemStoreInstance
}

func newInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		store:   make(map[string]VersionedValue),
		version: 0,
	}
}

func (s *InMemoryStore) BeginTransaction() *Transaction {
	return &Transaction{
		readSet:  make(map[string]VersionedValue),
		writeSet: make(map[string]string),
		delSet:   make(map[string]bool),
		ctx:      context.WithValue(context.Background(), "store", s),
	}
}

func (s *InMemoryStore) Get(key string) (string, error) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	val, ok := s.store[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return val.Val, nil
}

func (s *InMemoryStore) Set(key string, val string) error {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	versionVal, ok := s.store[key]
	// check if there are any verions of this key before,
	if !ok {
		vv := VersionedValue{
			Version: 1,
			Val:     val,
		}
		s.store[key] = vv
		return nil
	}
	// else, retrieve the latest versionedval, update both version + 1 and val in versionedval
	versionVal = VersionedValue{
		Version: versionVal.Version + 1,
		Val:     val,
	}

	// update store
	s.store[key] = versionVal

	return nil
}

func (s *InMemoryStore) Del(key string) error {

	_, ok := s.store[key]
	if !ok {
		return ErrKeyNotFound
	}

	delete(s.store, key)
	s.version++

	return nil
}

func (s *InMemoryStore) KeyVersion(key string) int {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	versionedVal, ok := s.store[key]

	if ok {
		return versionedVal.Version
	}
	return 0
}
