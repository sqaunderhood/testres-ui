This is a basic [TAP](https://testanything.org/) parser for the Go programming
language.

## Installation

Install or update using the `go get` command:

	go get -u github.com/ligurio/recidive/formats/testanything

## Example:

```go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ligurio/recidive/formats/testanything"
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

	_, err = NewParser(file)
	if err != nil {
		fmt.Println("Fail to parse")
	}
}
```
