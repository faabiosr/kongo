# Kongo
[![Build Status](https://img.shields.io/travis/fabiorphp/kongo/master.svg?style=flat-square)](https://travis-ci.org/fabiorphp/kongo)
[![Coverage Status](https://img.shields.io/coveralls/fabiorphp/kongo/master.svg?style=flat-square)](https://coveralls.io/github/fabiorphp/kongo?branch=master)

[Kong](https://getkong.org) Api Library for Golang

## Installation
Kongo requires Go 1.5 or later.
```
go get github.com/fabiorphp/kongo
```

## Usage
```go
package main

import (
    "github.com/fabiorphp/kongo"
)

func main() {
    kongo := kongo.New("127.0.0.1:8001")
    status := kongo.Node.Status()
    ...
}
```

## Full docs, see:
[https://godoc.org/github.com/fabiorphp/kongo](https://godoc.org/github.com/fabiorphp/kongo)

## License
This project is released under the MIT licence. See [LICENCE](https://github.com/fabiorphp/kongo/blob/master/LICENSE) for more details.
