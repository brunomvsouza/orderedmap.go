// generated by go run internal/gen.go -naturalType *uint; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap

import "sync"

// NewUintPtrMap return a new Map implemented by UintPtrMap
func NewUintPtrMap() *UintPtrMap {
	return &UintPtrMap{
		store: map[*uint]interface{}{},
		keys:  []*uint{},
	}
}

// UintPtrMap insertion ordered Map implementation
type UintPtrMap struct {
	sync.Mutex

	store map[*uint]interface{}
	keys  []*uint
}

// Put add a key-value pair to the UintPtrMap
func (m *UintPtrMap) Put(key *uint, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the UintPtrMap
func (m *UintPtrMap) Get(key *uint) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return
}

// Remove remove a key-value pair from the UintPtrMap
func (m *UintPtrMap) Remove(key *uint) {
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

// Empty return if the UintPtrMap in empty or not
func (m *UintPtrMap) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the UintPtrMap in insertion order
func (m *UintPtrMap) Keys() []*uint {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the UintPtrMap in insertion order
func (m *UintPtrMap) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the UintPtrMap
func (m *UintPtrMap) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
