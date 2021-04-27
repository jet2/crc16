
# crc16
[![Coverage Status](https://coveralls.io/repos/r10r/crc16/badge.svg?branch=master&service=github)](https://coveralls.io/github/jet2/crc16?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/jet2/crc16.svg)](https://pkg.go.dev/github.com/jet2/crc16)

Go implementation of CRC-16 calculation for majority of widely-used polynomials.</br>
It implements the golang hash.Hash interface.

## Usage

```go
package main

import (
	"fmt"

	"github.com/jet2/crc16"
)

func main() {
	table := crc16.MakeTable(crc16.CRC16_DF1)

	crc := crc16.Checksum([]byte("Hello world!"), table)
	fmt.Printf("CRC-16 DF1: %X\n", crc)

	// using the standard library hash.Hash interface
	h := crc16.New(table)
	h.Write([]byte("Hello world!"))
	fmt.Printf("CRC-16 DF1: %X\n", h.Sum16())
}
```
