package store

import "errors"

type CacheError error

var (
	ErrKeyNotFound      CacheError = errors.New("key not found")
	ErrCapacityExceeded CacheError = errors.New("capacity exceeded")
	ErrEviction         CacheError = errors.New("unable to evict")
)
