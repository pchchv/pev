# pev [![](https://godoc.org/github.com/pchchv/pev?status.svg)](https://godoc.org/github.com/pchchv/pev)

A simple password validator using raw entropy values.

## Installation

```sh
go get github.com/pchchv/pev
```

## Usage

```go
package main

import "github.com/wagslane/pev"

func main() {
    const minEntropyBits = 60

    // entropy is a float64, representing the strength in base 2 (bits)
    entropy := pev.GetEntropy("a longer password")

    // if password has enough entropy,
    // err is nil otherwise,
    // formatted error message is provided explaining how
    // to increase the strength of the password
    // (safe to show to the client)
    err := pev.Validate("some password", minEntropyBits)
}
```