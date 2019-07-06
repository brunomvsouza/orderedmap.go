# Insertion Ordered Map

[![Pipeline status](https://lab.bmvs.io/bs/orderedmap.go/badges/master/pipeline.svg)](https://lab.bmvs.io/bs/orderedmap.go/commits/master) [![Pipeline status](https://ci.appveyor.com/api/projects/status/i48873sk7aej379y/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/orderedmap-go-y9194/branch/master) [![Coverage report](https://lab.bmvs.io/bs/orderedmap.go/badges/master/coverage.svg)](https://lab.bmvs.io/bs/orderedmap.go/commits/master) [![Go Report Card](https://goreportcard.com/badge/lab.bmvs.io/bs/orderedmap.go)](https://goreportcard.com/report/lab.bmvs.io/bs/orderedmap.go) [![GoDoc Reference](https://godoc.org/go.bmvs.io/orderedmap?status.svg)](http://godoc.org/go.bmvs.io/orderedmap)

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
