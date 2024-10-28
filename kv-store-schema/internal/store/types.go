package store

import "errors"

type Types string

const (
	Int    Types = "int"
	Bool   Types = "bool"
	Float  Types = "float"
	String Types = "string"
)

type AttributeValue struct {
	Type Types
	Val  any
}

type Attribute struct {
	Key   string
	Value AttributeValue
}

type AttributeMap map[string]Attribute

type RegistryKey struct {
	Key     string
	AttrKey string
}

var (
	ErrKeyNotFound             error = errors.New("key not found")
	ErrInvalidAttrType               = errors.New("invalid attribute type")
	ErrSchemaAlreadyRegistered       = errors.New("schema already registered")
	ErrSchemaEntryNotFound           = errors.New("schema entry not found")
	ErrInvalidNumAttributes          = errors.New("invalid number of attributes")
)
