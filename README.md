# Insertion Ordered Map

[![GoDoc Reference](https://godoc.org/go.bmvs.io/orderedmap?status.svg)](http://godoc.org/go.bmvs.io/orderedmap)

Implementation of an insertion ordered map in Go

## Installation

```
go get go.bmvs.io/orderedmap
```

## Usage

```go
package main

import (
	"fmt"

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
