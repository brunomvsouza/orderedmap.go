// generated by go run internal/gen.go -naturalType uint64; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap

import "sync"

// NewUint64Map return a new Map implemented by Uint64Map
func NewUint64Map() *Uint64Map {
	return &Uint64Map{
		store: map[uint64]interface{}{},
		keys:  []uint64{},
	}
}

// Uint64Map insertion ordered Map implementation
type Uint64Map struct {
	sync.Mutex

	store map[uint64]interface{}
	keys  []uint64
}

// Put add a key-value pair to the Uint64Map
func (m *Uint64Map) Put(key uint64, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the Uint64Map
func (m *Uint64Map) Get(key uint64) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the Uint64Map
func (m *Uint64Map) Remove(key uint64) {
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

// Empty return if the Uint64Map in empty or not
func (m *Uint64Map) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the Uint64Map in insertion order
func (m *Uint64Map) Keys() []uint64 {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the Uint64Map in insertion order
func (m *Uint64Map) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the Uint64Map
func (m *Uint64Map) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
