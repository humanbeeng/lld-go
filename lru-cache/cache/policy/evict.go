package policy

type Evictor[K comparable] interface {
	UpdateKeyAccess(key K) error
	Evict() (K, error)
}
