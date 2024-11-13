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

## What value of entropy to use?
That's up to you. But a range of 50-70 seems “reasonable”.

## Workflow
First, a “base” number is determined. The base is the sum of the different “character sets” found in the password.
*The following character sets are chosen:
* 26 lowercase letters
* 26 uppercase letters
* 10 digits
* 5 substitution characters - `!@$&*'.
* 5 delimiter characters - `_-., `.
* 22 less common special characters - `'#%'()+/:;<=>?[\]^{|}~`.

If you use at least one character from each set, your base number will be 94: ``26+26+10+5+5+5+5+22 = 94``.

Each unique character that does not match any of these sets will add `1` to the base.

If you use only lowercase letters and numbers, for example, your base will be 36: `26+10 = 36`.

Once we have calculated the base, the total number of guesses for the bruteforce can be found using the following formula: `base^length`.

A base 26 password with 7 characters would require `26^7`, or `80318.
