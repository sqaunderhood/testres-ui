package subunit

import (
	"os"
	"testing"
)

type tcase struct {
	name string
	skip bool
}

var testset = []tcase{

	// OpenStack, SubUnit v2
	{name: "subunit-sample-01.subunit", skip: false},
	// OpenStack, SubUnit v2
	{name: "subunit-sample-02.subunit", skip: false},
	// Bazaar project, SubUnit v1
	{name: "subunit-sample-03.subunit", skip: false},
	// Bazaar project, SubUnit v1
	{name: "subunit-sample-04.subunit", skip: false},
}

func TestParser(t *testing.T) {
	for _, tcase := range testset {
		if tcase.skip {
			t.Logf("Skip: %s", tcase.name)
			continue
		}
		t.Logf("Running: %s", tcase.name)

		file, err := os.Open("tests/" + tcase.name)
		if err != nil {
			t.Fatal(err)
		}

		report, err := NewParser(file)
		if err != nil {
			t.Fatalf("Error parsing %s: %s", tcase.name, err)
		}

		if report == nil {
			t.Fatalf("Report == nil")
		}
	}
}
