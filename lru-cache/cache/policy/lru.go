package policy

import (
	"container/list"

	"github.com/humanbeeng/lld-go/lru-cache/cache/store"
)

type LRUEvictor[K comparable] struct {
	nodeMap map[K]*list.Element
	list    *list.List
}

func NewLRUEvictor[K comparable]() LRUEvictor[K] {
	return LRUEvictor[K]{
		nodeMap: make(map[K]*list.Element),
		list:    list.New(),
	}
}

func (e *LRUEvictor[K]) UpdateKeyAccess(key K) error {

	ele, ok := e.nodeMap[key]
	if ok {
		// Take the element and move to top
		e.list.MoveToFront(ele)
		return nil
	}

	// Create a new list element, update the nodeMap and insert to front of list
	ele = e.list.PushFront(key)
	e.nodeMap[key] = ele

	return nil
}

func (e *LRUEvictor[K]) Evict() (K, error) {
	var v K

	back := e.list.Back()
	if back == nil {
		return v, store.ErrEviction
	}
	e.list.Remove(back)

	delete(e.nodeMap, back.Value.(K))

	return back.Value.(K), nil

}
