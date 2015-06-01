package cowmap

import (
	"sync"
	"sync/atomic"
)

type Key interface{}

type Value interface{}

type Map struct {
	v atomic.Value
	l sync.Mutex
}

func New(m map[Key]Value) *Map {
	ret := new(Map)
	ret.v.Store(m)
	return ret
}

func (m *Map) Set(key Key, value Value) {
	m.l.Lock()
	defer m.l.Unlock()
	cur := m.v.Load().(map[Key]Value)
	newMap := make(map[Key]Value)
	for k, v := range cur {
		newMap[k] = v
	}
	newMap[key] = value
	m.v.Store(newMap)
}

func (m *Map) Get(key Key) Value {
	cur := m.v.Load().(map[Key]Value)
	return cur[key]
}

func (m *Map) Get2(key Key) (value Value, ok bool) {
	cur := m.v.Load().(map[Key]Value)
	value, ok = cur[key]
	return
}

func (m *Map) Del(key Key) {
	m.l.Lock()
	defer m.l.Unlock()
	cur := m.v.Load().(map[Key]Value)
	newMap := make(map[Key]Value)
	for k, v := range cur {
		newMap[k] = v
	}
	delete(newMap, key)
	m.v.Store(newMap)
}
