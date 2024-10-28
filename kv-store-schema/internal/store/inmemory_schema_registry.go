package store

import (
	"errors"
	"fmt"
)

type InMemorySchemaRegistry struct {
	registry map[RegistryKey]Types
}

func NewInMemorySchemaRegistry() InMemorySchemaRegistry {
	return InMemorySchemaRegistry{
		registry: make(map[RegistryKey]Types),
	}
}

func (sr *InMemorySchemaRegistry) Register(key RegistryKey, value any) (Types, error) {
	if sr.Exists(key) {
		return "", errors.Join(ErrSchemaAlreadyRegistered, fmt.Errorf("key %v-%v", key.Key, key.AttrKey))
	}

	valType, err := getValueType(value)
	if err != nil {
		return "", err
	}

	sr.registry[key] = valType

	return valType, nil
}

func (sr *InMemorySchemaRegistry) GetRegisteredType(key RegistryKey) (Types, error) {
	if !sr.Exists(key) {
		return "", errors.Join(ErrKeyNotFound, fmt.Errorf("key: %v-%v", key.Key, key.AttrKey))
	}
	return sr.registry[key], nil
}

func (sr *InMemorySchemaRegistry) Validate(key RegistryKey, value any) error {
	if !sr.Exists(key) {
		return nil
	}

	typ, err := sr.GetRegisteredType(key)
	if err != nil {
		return err
	}

	valType, err := getValueType(value)
	if err != nil {
		return err
	}

	if valType != typ {
		return errors.Join(ErrInvalidAttrType, fmt.Errorf("expected %v, given %v", typ, valType))
	}

	return nil
}

func (sr *InMemorySchemaRegistry) Exists(key RegistryKey) bool {
	_, ok := sr.registry[key]
	return ok
}

func (sr *InMemorySchemaRegistry) Unregister(key RegistryKey) error {
	if !sr.Exists(key) {
		return errors.Join(ErrSchemaEntryNotFound, fmt.Errorf("key: %v-%v", key.Key, key.AttrKey))
	}

	delete(sr.registry, key)

	return nil
}
