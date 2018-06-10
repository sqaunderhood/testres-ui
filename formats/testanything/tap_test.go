package tap

import (
	"os"
	"testing"
)

type tcase struct {
	name string
	skip bool
}

var testset = []tcase{

	{name: "tap-sample-01.tap", skip: false},
	{name: "tap-sample-02.tap", skip: false},
	{name: "tap-sample-03.tap", skip: false},
	{name: "tap-sample-04.tap", skip: false},
	{name: "tap-sample-05.tap", skip: false},
	{name: "tap-sample-06.tap", skip: false},
	{name: "tap-sample-07.tap", skip: false},
	{name: "tap-sample-08.tap", skip: false},
	// CRIU tests
	{name: "tap-sample-09.tap", skip: false},
	// GIT tests
	{name: "tap-sample-10.tap", skip: false},
	// LibVirt tests (libvirt-tck-f10-broken.txt)
	{name: "tap-sample-11.tap", skip: false},
	// LibVirt tests (libvirt-tck-f10-fixed.txt)
	{name: "tap-sample-12.tap", skip: false},
	/*
	TODO:
	- pytest-tap https://github.com/python-tap/pytest-tap
	- postgresql project
	*/
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
			t.Fatalf("error parsing %s: %s", tcase.name, err)
		}

		suite, err := report.Suite()
		if err != nil {
			panic(err)
		} else if suite == nil {
			t.Fatalf("Suite == nil")
		}
	}
}
