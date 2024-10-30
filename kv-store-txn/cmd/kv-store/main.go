package main

import (
	"fmt"

	"github.com/humanbeeng/lld-go/kv-store-txn/internal/store"
)

func main() {

	fmt.Println("Transactional KV-Store")
	var st store.Store

	inMemStore := store.GetInMemoryStore()
	st = inMemStore

	// init set
	st.Set("name", "nithin")

	name, err := st.Get("name")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(name)

	// transaction begin
	tx1 := st.BeginTransaction()

	name, _ = tx1.Get(st, "name")
	fmt.Println("tx", "name", name)

	// some other thread updated. "nithin" -> "lol"
	// st.Del("name")

	// "nithin" -> "raju"
	tx1.Set(st, "name", "Raju")

	name, _ = st.Get("name") // returns "nithin"
	fmt.Println("st", "name", name)

	tx1.Del(st, "name")
	// should throw error
	err = tx1.Commit(st)
	fmt.Println("commit")

	if err != nil {
		fmt.Println("err while commit", err)
		return
	}

	name, _ = st.Get("name")
	fmt.Println("st", "name", name)
}
