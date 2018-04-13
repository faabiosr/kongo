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

- Install [go dep](https://github.com/golang/dep)

### Run tests
```sh
// tests
$ make test

// test with coverage
$ make test-coverage

// clean-up
$ make clean

// configure (download dependencies)
$ make configure
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/fabiorphp/kongo/blob/master/LICENSE) for more details.
