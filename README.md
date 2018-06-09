# Kongo

[![Build Status](https://img.shields.io/travis/fabiorphp/kongo/master.svg?style=flat-square)](https://travis-ci.org/fabiorphp/kongo)
[![Coverage Status](https://img.shields.io/coveralls/fabiorphp/kongo/master.svg?style=flat-square)](https://coveralls.io/github/fabiorphp/kongo?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/fabiorphp/kongo)
[![Go Report Card](https://goreportcard.com/badge/github.com/fabiorphp/kongo?style=flat-square)](https://goreportcard.com/report/github.com/fabiorphp/kongo)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/fabiorphp/kongo/blob/master/LICENSE)

[Kong](https://getkong.org) api library for Golang

## Installation

Kongo requires Go 1.9 or later.

```
go get github.com/fabiorphp/kongo
```

If you want to get an specific version, please use the example below:

```
go get gopkg.in/fabiorphp/kongo.v0
```

## Usage
```go
package main

import (
    "github.com/fabiorphp/kongo"
)

func main() {
    kongo := kongo.New(nil, "127.0.0.1:8001")
    status, _, _ := kongo.Node.Status()
    ...
}
```

## Documentation

Read the full documentation at [https://godoc.org/github.com/fabiorphp/kongo](https://godoc.org/github.com/fabiorphp/kongo).

## Development

### Requirements

- Install [Go](https://golang.org)
- Install [go dep](https://github.com/golang/dep)

### Makefile
```sh
// Clean up
$ make clean

// Creates folders and download dependencies
$ make configure

//Run tests and generates html coverage file
make cover

// Download project dependencies
make depend

// Format all go files
make fmt

//Run linters
make lint

// Run tests
make test
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/fabiorphp/kongo/blob/master/LICENSE) for more details.
