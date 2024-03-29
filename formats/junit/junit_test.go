package junit

import (
	"os"
	"testing"
)

type tcase struct {
	name   string  // name of file
	skip   bool    // flag to skip
	suites []suite //
}

type suite struct {
	tc_number   int // number of testcases
	prop_number int // number of properties
}

var testset = []tcase{
	{
		name:   "junit-sample-1.xml",
		skip:   false,
		suites: []suite{{tc_number: 8, prop_number: 0}},
	},
	{
		name:   "junit-sample-2.xml",
		skip:   false,
		suites: []suite{{tc_number: 7, prop_number: 0}},
	},
	{
		name:   "junit-sample-3.xml",
		skip:   false,
		suites: []suite{{tc_number: 2, prop_number: 0}, {tc_number: 2, prop_number: 0}},
	},
	{
		name: "junit-sample-4.xml",
		skip: false,

		suites: []suite{{tc_number: 9, prop_number: 0}, {tc_number: 7, prop_number: 0}},
	},
	{
		name: "junit-sample-5.xml",
		skip: false,
		suites: []suite{

			{tc_number: 16, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 13, prop_number: 0},
			{tc_number: 25, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 20, prop_number: 0},
			{tc_number: 41, prop_number: 0},
			{tc_number: 17, prop_number: 0},
			{tc_number: 14, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 17, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 19, prop_number: 0},
			{tc_number: 5, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 29, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 17, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 16, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 31, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 39, prop_number: 0},
			{tc_number: 18, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 42, prop_number: 0},
			{tc_number: 76, prop_number: 0},
			{tc_number: 15, prop_number: 0},
			{tc_number: 24, prop_number: 0},
			{tc_number: 14, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 24, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 13, prop_number: 0},
			{tc_number: 16, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 29, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 23, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 19, prop_number: 0},
			{tc_number: 16, prop_number: 0},
			{tc_number: 17, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 19, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 21, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 27, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 46, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 19, prop_number: 0},
			{tc_number: 53, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 11, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 22, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 19, prop_number: 0},
			{tc_number: 22, prop_number: 0},
			{tc_number: 26, prop_number: 0},
			{tc_number: 5, prop_number: 0},
			{tc_number: 22, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 5, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 13, prop_number: 0},
			{tc_number: 19, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 34, prop_number: 0},
			{tc_number: 23, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 20, prop_number: 0},
			{tc_number: 5, prop_number: 0},
			{tc_number: 45, prop_number: 0},
			{tc_number: 14, prop_number: 0},
			{tc_number: 16, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 58, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 26, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 13, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 14, prop_number: 0},
			{tc_number: 13, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 60, prop_number: 0},
			{tc_number: 5, prop_number: 0},
			{tc_number: 31, prop_number: 0},
			{tc_number: 47, prop_number: 0},
			{tc_number: 12, prop_number: 0},
			{tc_number: 5, prop_number: 0},
			{tc_number: 22, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 22, prop_number: 0},
			{tc_number: 34, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 8, prop_number: 0},
			{tc_number: 10, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 38, prop_number: 0},
			{tc_number: 34, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 26, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 7, prop_number: 0},
			{tc_number: 15, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 6, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 37, prop_number: 0},
			{tc_number: 2, prop_number: 0},
			{tc_number: 9, prop_number: 0},
			{tc_number: 7, prop_number: 0},
		},
	},
	{
		name: "junit-sample-6.xml",
		skip: false,
		suites: []suite{{tc_number: 1, prop_number: 0},
			{tc_number: 1, prop_number: 0},
			{tc_number: 1, prop_number: 0},
			{tc_number: 1, prop_number: 0},
			{tc_number: 1, prop_number: 0},
			{tc_number: 1, prop_number: 0},
		},
	},
	{
		/*
			Sample of JUnit report, that is successfully parsed by Bamboo
			https://confluence.atlassian.com/bamboo/junit-parsing-in-bamboo-289277357.html
		*/
		name: "junit-sample-7.xml",
		skip: true,
	},
	{
		/*
			Sample of JUnit report, that is successfully parsed by Bamboo
			https://confluence.atlassian.com/bamboo/junit-parsing-in-bamboo-289277357.html
		*/
		name: "junit-sample-8.xml",
		skip: true,
	},
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

		if len(report.Suites) != len(tcase.suites) {
			t.Fatalf("Wrong number of suites: %s", tcase.name)
		}

		for num, s := range report.Suites {
			if len(s.TestCases) != tcase.suites[num].tc_number {
				t.Fatalf("Wrong number of testcases: %s", tcase.name)
			}
			if len(s.Properties) != tcase.suites[num].prop_number {
				t.Fatalf("Wrong number of properties: %s", tcase.name)
			}
		}
	}
}
