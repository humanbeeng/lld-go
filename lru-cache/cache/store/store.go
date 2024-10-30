package store

type Store[K comparable, V comparable] interface {
	Get(key K) (V, error)
	Set(key K, value V) error
	Del(key K) error
}
