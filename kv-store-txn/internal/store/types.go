package store

import "errors"

type VersionedValue struct {
	Version int
	Val     string
}

type StoreErr error

var (
	ErrKeyNotFound StoreErr = errors.New("key not found")
)
