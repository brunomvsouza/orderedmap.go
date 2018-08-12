# Insertion Ordered Map

[![GoDoc Reference](https://godoc.org/go.bmvs.io/orderedmap?status.svg)](http://godoc.org/go.bmvs.io/orderedmap) [![Build Status](https://travis-ci.com/brunomvsouza/orderedmap.go.svg?branch=master)](https://travis-ci.com/brunomvsouza/orderedmap.go) [![Build status](https://ci.appveyor.com/api/projects/status/m7c7tlk6bdfh3p35/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/orderedmap-go/branch/master) [![Coverage Status](https://coveralls.io/repos/github/brunomvsouza/orderedmap.go/badge.svg?branch=master)](https://coveralls.io/github/brunomvsouza/orderedmap.go)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedmap.go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedmap.go?ref=badge_shield)

Implementation of an insertion ordered map in Go

## Installation

```
go get go.bmvs.io/orderedmap
```

## Usage

```go
package main

import (
	"go.bmvs.io/orderedmap"
)

func main() {
	m := orderedmap.New()
	m.Put("foo", "bar")
	// ...
}
```

See the [godoc](https://godoc.org/go.bmvs.io/orderedmap) to see all the available methods with example usage.

## Development

- Run tests with `go test -race ./...`

## License

BSD-2-Clause


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedmap.go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedmap.go?ref=badge_large)