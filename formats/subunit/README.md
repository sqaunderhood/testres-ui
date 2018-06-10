This is a basic subunit parser for the Go programming language. It supports
reading of test reports formatted in
[SubUnit](https://github.com/testing-cabal/subunit) format (both versions v1
and v2 are supported).

## Installation

Install or update using the `go get` command:

	go get -u github.com/ligurio/go-subunit

## Example

```go
package main

import (
        "flag"
        "fmt"
        . "github.com/ligurio/go-subunit/parser"
        "os"
)

func main() {

        var f = flag.String("file", "", "filename")
        flag.Parse()

        file, err := os.Open(*f)
        if err != nil {
                fmt.Println("Error opening file:", err)
                return
        }
        defer file.Close()

        err, t := Parser(file)

        fmt.Printf("%#v\n", t)
}
```
