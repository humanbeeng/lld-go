package cache

import (
	"github.com/humanbeeng/lld-go/lru-cache/cache/policy"
	"github.com/humanbeeng/lld-go/lru-cache/cache/store"
)

type Cache[K comparable, V comparable] struct {
	store   store.Store[K, V]
	evictor policy.Evictor[K]
}

func NewCache[K, V comparable](store store.Store[K, V], evictor policy.Evictor[K]) Cache[K, V] {
	return Cache[K, V]{
		store:   store,
		evictor: evictor,
	}
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	val, err := c.store.Get(key)
	if err != nil {
		return val, err
	}

	err = c.evictor.UpdateKeyAccess(key)
	if err != nil {
		return val, err
	}

	return val, nil
}

func (c *Cache[K, V]) Set(key K, value V) error {
	// check for capacity
	err := c.store.Set(key, value)
	if err != nil {
		// call evict
		evictedKey, err := c.evictor.Evict()
		if err != nil {
			return err
		}
		c.Del(evictedKey)

		err = c.store.Set(key, value)
		if err != nil {
			return err
		}
		return nil
	}

	err = c.evictor.UpdateKeyAccess(key)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache[K, V]) Del(key K) error {
	err := c.store.Del(key)
	if err != nil {
		return err
	}

	return nil
}
