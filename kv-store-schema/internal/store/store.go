package store

type KVStore interface {
	Get(key string) (AttributeMap, error)
	GetAttr(key string, attrKey string) (Attribute, error)
	Put(key string, attrs ...Attribute) error
	Delete(key string) error
	DeleteAttr(key string, attrKey string) error
}
