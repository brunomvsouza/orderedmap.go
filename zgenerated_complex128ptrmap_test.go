// generated by go run internal/gen.go -naturalType *complex128; DO NOT EDIT.

// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedmap_test

import (
	"testing"

	"go.bmvs.io/orderedmap"
)

func TestComplex128PtrMap_Put(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

    m := orderedmap.NewComplex128PtrMap()
	m.Put(e1, 1)
	m.Put(e2, 1)

	if m.Size() != 2 {
		t.Errorf("expected 2 entries, got %d", m.Size())
	}
}

func TestComplex128PtrMap_Get(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

    m := orderedmap.NewComplex128PtrMap()
	m.Put(e1, 1)
	m.Put(e2, 1)

	value, found := m.Get(e1)
	if !found {
	    t.Errorf("expected true, got %s", "false")
	}
	if value != 1 {
	    t.Errorf("expected 1, got %+v", value)
	}
}

func TestComplex128PtrMap_Remove(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

	m := orderedmap.NewComplex128PtrMap()
    m.Put(e1, 1)
    m.Put(e2, 1)

	m.Remove(e1)

	if m.Size() != 1 {
        t.Errorf("expected 1 entries, got %d", m.Size())
    }
}

func TestComplex128PtrMap_Size(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

	m := orderedmap.NewComplex128PtrMap()
    m.Put(e1, 1)
    m.Put(e2, 1)

	if m.Size() != 2 {
		t.Errorf("expected the set size to be 2 but it was %d", m.Size())
	}
}

func TestComplex128PtrMap_Empty(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

	m := orderedmap.NewComplex128PtrMap()
	if !m.Empty() {
		t.Errorf("expected new set to be empty but it had %d elements", m.Size())
	}

	m.Put(e1, 1)
    m.Put(e2, 1)

	if m.Empty() {
		t.Error("expected a set with added items to not be empty")
	}
}

func TestComplex128PtrMap_Keys(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

	m := orderedmap.NewComplex128PtrMap()
    m.Put(e1, 1)
    m.Put(e2, 1)

	keys := m.Keys()
	if keys[0] != e1 {
	    t.Errorf("expected the key %d to be %+v", 0, e1)
	}
	if keys[1] != e2 {
        t.Errorf("expected the key %d to be %+v", 0, e2)
    }
}

func TestComplex128PtrMap_Values(t *testing.T) {
	var e1, e2 *complex128
	e := createRandomObject(e1)
	if v, ok := e.(*complex128); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(*complex128); ok {
		e2 = v
	}

	m := orderedmap.NewComplex128PtrMap()
    m.Put(e1, 1)
    m.Put(e2, 1)

	values := m.Values()
	if values[0] != 1 {
	    t.Errorf("expected the key %d to be %+v", 0, 1)
	}
	if values[1] != 1 {
        t.Errorf("expected the key %d to be %+v", 0, 1)
    }
}
