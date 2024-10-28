package store

type SchemaRegistry interface {
	Register(key RegistryKey, value any) error
	Validate(key RegistryKey, value any) error

	GetRegisteredType(key RegistryKey) (Types, error)
	Exists(key RegistryKey) bool
	Unregister(key RegistryKey) error
}
