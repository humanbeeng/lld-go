package store

import (
	"errors"
	"fmt"
	"sync"
)

type InMemoryStore struct {
	store          map[string]AttributeMap
	schemaRegistry InMemorySchemaRegistry
	sync.RWMutex
}

func (s *InMemoryStore) Get(key string) (AttributeMap, error) {
	attrMap, ok := s.store[key]
	if !ok {
		return attrMap, errors.Join(ErrKeyNotFound, fmt.Errorf("key: %v", key))
	}

	return attrMap, nil
}

func (s *InMemoryStore) GetAttr(key, attrKey string) (Attribute, error) {
	attrMap, ok := s.store[key]
	if !ok {
		return Attribute{}, errors.Join(ErrKeyNotFound, fmt.Errorf("key: %v", key))
	}

	attr, ok := attrMap[attrKey]
	if !ok {
		return Attribute{}, errors.Join(ErrKeyNotFound, fmt.Errorf("attribute key: %v", attrKey))
	}

	return attr, nil
}

func (s *InMemoryStore) Put(key string, attrs ...Attribute) error {
	// check if key exists
	// if not, then register schema. construct an attributeMap. insert into key. return

	// if key exists. get the attributeMap
	// for each attrs, check if there is an extry in schemaregistry, if yes, validate the schema. and update
	// else, register schema, insert into attributeMap and insert into store

	return nil
}

func (s *InMemoryStore) Delete(key string) error {
	return nil
}

func (s *InMemoryStore) DeleteAttr(key, attrKey string) error {
	return nil
}
