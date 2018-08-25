// generated by go run internal/gen.go -naturalType *int32; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap

import "sync"

// NewInt32PtrMap return a new Map implemented by Int32PtrMap
func NewInt32PtrMap() *Int32PtrMap {
	return &Int32PtrMap{
		store: map[*int32]interface{}{},
		keys:  []*int32{},
	}
}

// Int32PtrMap insertion ordered Map implementation
type Int32PtrMap struct {
	sync.Mutex

	store map[*int32]interface{}
	keys  []*int32
}

// Put add a key-value pair to the Int32PtrMap
func (m *Int32PtrMap) Put(key *int32, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the Int32PtrMap
func (m *Int32PtrMap) Get(key *int32) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the Int32PtrMap
func (m *Int32PtrMap) Remove(key *int32) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		return
	}

	delete(m.store, key)

	for i, _ := range m.keys {
		if m.keys[i] == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

// Empty return if the Int32PtrMap in empty or not
func (m *Int32PtrMap) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the Int32PtrMap in insertion order
func (m *Int32PtrMap) Keys() []*int32 {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the Int32PtrMap in insertion order
func (m *Int32PtrMap) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the Int32PtrMap
func (m *Int32PtrMap) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
