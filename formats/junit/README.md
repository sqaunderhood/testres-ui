This is a basic JUnit parser for the Go programming language.

As there is no definitive JUnit XSD that I could find, the XML documents read by
this package should correspond to a [JUnit XSD
scheme](https://svn.jenkins-ci.org/trunk/hudson/dtkit/dtkit-format/dtkit-junit-model/src/main/resources/com/thalesgroup/dtkit/junit/model/xsd/).
File a bug if something doesn't work like you expect it to.

## Installation

Install or update using the `go get` command:

	go get -u github.com/ligurio/recidive/junit

## Example

```go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ligurio/go-junit/"
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

	report, err := junit.NewParser(file)
	if err != nil {
		fmt.Printf("Failed to parse")
	}

	fmt.Println("JUnit report", *f)
	if report.Name != "" {
		fmt.Println("NAME", report.Name)
	}
	fmt.Println("Passed", report.Tests-report.Failures-report.Disabled-report.Errors)
	fmt.Println("Failed", report.Failures)
	fmt.Println("Total testsuites:", len(report.Suites), "\n")

	for i, s := range report.Suites {
		fmt.Println("Suite #", i, " - (", len(s.TestCases), " cases)")
		fmt.Println("\tSuite name:", s.Name)
		fmt.Println("\tTotal time:", s.Time)
		fmt.Println("\tFailures", s.Failures, "Passed", s.Tests-s.Disabled-s.Skipped-s.Errors)
	}

	fmt.Printf("%#v", report)
}
```
