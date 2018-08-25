// generated by go run internal/gen.go -naturalType *int64; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap

import "sync"

// NewInt64PtrMap return a new Map implemented by Int64PtrMap
func NewInt64PtrMap() *Int64PtrMap {
	return &Int64PtrMap{
		store: map[*int64]interface{}{},
		keys:  []*int64{},
	}
}

// Int64PtrMap insertion ordered Map implementation
type Int64PtrMap struct {
	sync.Mutex

	store map[*int64]interface{}
	keys  []*int64
}

// Put add a key-value pair to the Int64PtrMap
func (m *Int64PtrMap) Put(key *int64, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the Int64PtrMap
func (m *Int64PtrMap) Get(key *int64) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the Int64PtrMap
func (m *Int64PtrMap) Remove(key *int64) {
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

// Empty return if the Int64PtrMap in empty or not
func (m *Int64PtrMap) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the Int64PtrMap in insertion order
func (m *Int64PtrMap) Keys() []*int64 {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the Int64PtrMap in insertion order
func (m *Int64PtrMap) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the Int64PtrMap
func (m *Int64PtrMap) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
