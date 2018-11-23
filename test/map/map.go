package main

import (
	"sync"
)

func main() {
}

type Value struct {
	int64  Value
}
type Map struct {
	Data map[string]interface{}
	Lock sync.RWMutex
}

func (m *Map) Set(k string, v interface{}) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	m.Data[k] = v
}
func (m *Map) Get(k string) interface{} {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	return m.Data[k]
}
func (m *Map) Delete(k string){
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	delete(m.Data,k)
}