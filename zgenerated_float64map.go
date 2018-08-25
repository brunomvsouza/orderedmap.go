// generated by go run internal/gen.go -naturalType float64; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap

import "sync"

// NewFloat64Map return a new Map implemented by Float64Map
func NewFloat64Map() *Float64Map {
	return &Float64Map{
		store: map[float64]interface{}{},
		keys:  []float64{},
	}
}

// Float64Map insertion ordered Map implementation
type Float64Map struct {
	sync.Mutex

	store map[float64]interface{}
	keys  []float64
}

// Put add a key-value pair to the Float64Map
func (m *Float64Map) Put(key float64, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the Float64Map
func (m *Float64Map) Get(key float64) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the Float64Map
func (m *Float64Map) Remove(key float64) {
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

// Empty return if the Float64Map in empty or not
func (m *Float64Map) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the Float64Map in insertion order
func (m *Float64Map) Keys() []float64 {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the Float64Map in insertion order
func (m *Float64Map) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the Float64Map
func (m *Float64Map) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
