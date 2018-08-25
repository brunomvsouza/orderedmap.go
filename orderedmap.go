// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

// Package orderedmap implements a thread safe insertion ordered map.
package orderedmap // import "go.bmvs.io/orderedmap"

//go:generate go run internal/gen.go -naturalType interface{} -noTest true

//go:generate go run internal/gen.go -naturalType rune

//go:generate go run internal/gen.go -naturalType complex128
//go:generate go run internal/gen.go -naturalType complex64
//go:generate go run internal/gen.go -naturalType *complex128
//go:generate go run internal/gen.go -naturalType *complex64

//go:generate go run internal/gen.go -naturalType int
//go:generate go run internal/gen.go -naturalType int8
//go:generate go run internal/gen.go -naturalType int16
//go:generate go run internal/gen.go -naturalType int32
//go:generate go run internal/gen.go -naturalType int64
//go:generate go run internal/gen.go -naturalType *int
//go:generate go run internal/gen.go -naturalType *int8
//go:generate go run internal/gen.go -naturalType *int16
//go:generate go run internal/gen.go -naturalType *int32
//go:generate go run internal/gen.go -naturalType *int64

//go:generate go run internal/gen.go -naturalType uint
//go:generate go run internal/gen.go -naturalType uint8
//go:generate go run internal/gen.go -naturalType uint16
//go:generate go run internal/gen.go -naturalType uint32
//go:generate go run internal/gen.go -naturalType uint64
//go:generate go run internal/gen.go -naturalType uintptr
//go:generate go run internal/gen.go -naturalType *uint
//go:generate go run internal/gen.go -naturalType *uint8
//go:generate go run internal/gen.go -naturalType *uint16
//go:generate go run internal/gen.go -naturalType *uint32
//go:generate go run internal/gen.go -naturalType *uint64
//go:generate go run internal/gen.go -naturalType *uintptr

//go:generate go run internal/gen.go -naturalType float32
//go:generate go run internal/gen.go -naturalType float64
//go:generate go run internal/gen.go -naturalType *float32
//go:generate go run internal/gen.go -naturalType *float64
