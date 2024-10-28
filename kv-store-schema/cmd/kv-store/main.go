package main

import (
	"fmt"

	"github.com/humanbeeng/lld-go/kv-store-schema/internal/store"
)

func main() {
	fmt.Println("KV Store")

	var (
		registry store.SchemaRegistry
		kvstore  store.KVStore
	)

	inmemRegistry := store.NewInMemorySchemaRegistry()
	registry = &inmemRegistry

	inmemKvStore := store.NewInMemoryStore(registry)
	kvstore = &inmemKvStore

	age := store.Attribute{
		Key: "age",
		Value: store.AttributeValue{
			Val: 25,
		},
	}
	address := store.Attribute{
		Key: "address",
		Value: store.AttributeValue{
			Val: "Bangalore",
		},
	}

	err := kvstore.Put("nithin", age, address)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(kvstore.Get("nithin"))

	age.Value.Val = "24"
	err = kvstore.Put("nithin", age, address)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(kvstore.Get("nithin"))
}
