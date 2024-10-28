package store

import (
	"errors"
	"fmt"
	"sync"
)

type InMemoryStore struct {
	store          map[string]AttributeMap
	schemaRegistry SchemaRegistry
	sync.RWMutex
}

func NewInMemoryStore(schemaRegistry SchemaRegistry) InMemoryStore {
	return InMemoryStore{
		store:          make(map[string]AttributeMap),
		schemaRegistry: schemaRegistry,
	}
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

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	if len(attrs) == 0 {
		return ErrInvalidNumAttributes
	}

	attrMap, ok := s.store[key]
	// check if key exists
	// if not, then register schema. construct an attributeMap. insert into key. return
	if !ok {
		for _, attr := range attrs {
			registryKey := RegistryKey{Key: key, AttrKey: attr.Key}
			typ, err := s.schemaRegistry.Register(registryKey, attr.Value.Val)
			if err != nil {
				return err
			}
			attrMap = make(AttributeMap)

			attr.Value.Type = typ
			attrMap[key] = attr
		}

		s.store[key] = attrMap

		return nil
	}

	// if key exists. get the attributeMap
	// for each attrs, check if there is an extry in schemaregistry, if yes, validate the schema. and update
	for _, attr := range attrs {
		registryKey := RegistryKey{Key: key, AttrKey: attr.Key}

		if err := s.schemaRegistry.Validate(registryKey, attr.Value.Val); err != nil {
			return err
		}
	}

	// else, register schema, insert into attributeMap and insert into store
	for _, attr := range attrs {
		registryKey := RegistryKey{Key: key, AttrKey: attr.Key}

		if !s.schemaRegistry.Exists(registryKey) {
			s.schemaRegistry.Register(registryKey, attr.Value.Val)
			continue
		}

		attrMap[attr.Key] = attr
	}

	s.store[key] = attrMap

	return nil
}

func (s *InMemoryStore) Delete(key string) error {
	return nil
}

func (s *InMemoryStore) DeleteAttr(key, attrKey string) error {
	return nil
}
