// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

// Package orderedmap implements a thread safe insertion ordered map.
package orderedmap // import "go.bmvs.io/orderedmap"

import (
	"sync"
)

// New return a new Map implemented by OrderedMap
func New() *OrderedMap {
	return &OrderedMap{
		store: make(map[interface{}]interface{}),
		keys:  make([]interface{}, 0, 0),
	}
}

// OrderedMap insertion ordered Map implementation
type OrderedMap struct {
	sync.Mutex

	store map[interface{}]interface{}
	keys  []interface{}
}

// Put add a key-value pair to the OrderedMap
func (m *OrderedMap) Put(key, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the OrderedMap
func (m *OrderedMap) Get(key interface{}) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the OrderedMap
func (m *OrderedMap) Remove(key interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		return
	}

	delete(m.store, key)

	for i := range m.keys {
		if m.keys[i] == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

// Empty return if the OrderedMap in empty or not
func (m *OrderedMap) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the OrderedMap in insertion order
func (m *OrderedMap) Keys() []interface{} {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the OrderedMap in insertion order
func (m *OrderedMap) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the OrderedMap
func (m *OrderedMap) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
