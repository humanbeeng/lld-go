package store

import (
	"context"
	"fmt"
	"sync"
)

type Transaction struct {
	readSet  map[string]VersionedValue
	writeSet map[string]string
	delSet   map[string]bool

	sync.RWMutex
	ctx context.Context
}

func (t *Transaction) Get(store Store, key string) (string, error) {
	// check if there are any modificiations done in this transaction. check writeSet
	// if so, return the buffered val
	if val, ok := t.writeSet[key]; ok {
		return val, nil
	}

	// insert into readSet. key -> version
	val, err := store.Get(key)
	if err != nil {
		return "", err
	}

	version := store.KeyVersion(key)
	t.readSet[key] = VersionedValue{
		Version: version,
		Val:     val,
	}

	return val, nil
}

func (t *Transaction) Set(store Store, key string, val string) (string, error) {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	t.writeSet[key] = val
	return "", nil
}

func (t *Transaction) Del(store Store, key string) error {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	t.delSet[key] = true
	delete(t.writeSet, key)

	return nil
}

func (t *Transaction) Commit(store Store) error {
	// call validate
	if err := t.validate(store); err != nil {
		t.Rollback(store)
		return err
	}

	// acquire lock
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	// Set all the values present in writeSet
	for k, v := range t.writeSet {
		store.Set(k, v)
	}

	// del all keys in delSet
	for k := range t.delSet {
		store.Del(k)
	}

	return nil
}

func (t *Transaction) Rollback(store Store) error {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	t.readSet = nil
	t.writeSet = nil
	t.delSet = nil

	return nil
}

func (t *Transaction) validate(store Store) error {
	for k, v := range t.readSet {
		version := store.KeyVersion(k)
		fmt.Println("txn version", "v", "found version", version)
		if version != v.Version {
			return fmt.Errorf("Conflict detected")
		}
	}
	return nil
}
