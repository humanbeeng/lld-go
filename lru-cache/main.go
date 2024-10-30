package main

import (
	"fmt"

	"github.com/humanbeeng/lld-go/lru-cache/cache"
	"github.com/humanbeeng/lld-go/lru-cache/cache/policy"
	"github.com/humanbeeng/lld-go/lru-cache/cache/store"
)

func main() {
	var evictor policy.Evictor[string]
	lru := policy.NewLRUEvictor[string]()

	var st store.Store[string, string]
	inmem := store.GetInMemoryStore(3)

	st = inmem
	evictor = &lru

	c := cache.NewCache[string, string](st, evictor)

	err := c.Set("name", "nithin")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	name, err := c.Get("name")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("GET:", "name", name)

	err = c.Set("age", "24")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	err = c.Set("addr", "bangalore")
	if err != nil {
		fmt.Println(err)
		return
	}

	age, err := c.Get("age")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("GET:", "age", age)

	err = c.Set("image", "https")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	image, err := c.Get("image")
	fmt.Println(image)

}
