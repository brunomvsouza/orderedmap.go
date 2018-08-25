// generated by go run internal/gen.go -naturalType uint8; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap

import "sync"

// NewUint8Map return a new Map implemented by Uint8Map
func NewUint8Map() *Uint8Map {
	return &Uint8Map{
		store: map[uint8]interface{}{},
		keys:  []uint8{},
	}
}

// Uint8Map insertion ordered Map implementation
type Uint8Map struct {
	sync.Mutex

	store map[uint8]interface{}
	keys  []uint8
}

// Put add a key-value pair to the Uint8Map
func (m *Uint8Map) Put(key uint8, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the Uint8Map
func (m *Uint8Map) Get(key uint8) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the Uint8Map
func (m *Uint8Map) Remove(key uint8) {
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

// Empty return if the Uint8Map in empty or not
func (m *Uint8Map) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the Uint8Map in insertion order
func (m *Uint8Map) Keys() []uint8 {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the Uint8Map in insertion order
func (m *Uint8Map) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the Uint8Map
func (m *Uint8Map) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
