package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"

	"github.com/ligurio/go-junit/parser"
)

type Status int
type Format int

const (
	StatusNone Status = iota
	StatusFail
	StatusPass
	StatusSkip
	StatusError
	StatusTodo
	StatusXFail
	StatusUxSuccess
)

const (
	FmtSubUnit Format = iota
	FmtJUnit
	FmtTAP
)

var (
	// Statuses maps status to its friendly name
	Statuses = map[Status]string{
		StatusNone:      "NONE",
		StatusFail:      "FAIL",
		StatusPass:      "PASS",
		StatusSkip:      "SKIP",
		StatusError:     "ERROR",
		StatusTodo:      "TODO",
		StatusXFail:     "XFAIL",
		StatusUxSuccess: "UXSUCCESS",
	}
)

var (
	// Formats maps format to its friendly name
	Formats = map[Format]string{
		FmtSubUnit: "SubUnit",
		FmtJUnit:   "JUnit",
		FmtTAP:     "TAP (Test Anything Protocol)",
	}
)

func ReadReport(r io.Reader, name string) (*Report, error) {

	report := new(Report)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	report.Body = string(buf)

	r = bytes.NewReader(buf)

	report.UID = makeID()
	log.Println("Report ID is", report.UID)

	report.Filename = name

	if jreport, err := junit.NewParser(r); err == nil {
		log.Println("DEBUG: JUnit format detected.")
		report.Format = Formats[FmtJUnit]

		for _, ts := range jreport.Suites {
			s := new(Suite)
			for _, test := range ts.TestCases {
				t := new(Test)
				t.Name = test.Name
				//t.Status = test.Status
				//t.Ok =
				t.Description = ""
				t.Explanation = ""
				t.StartTime = test.Time
				t.EndTime = test.Time
				//t.Tags =
				s.Tests = append(s.Tests, *t)
			}
			report.Suites = append(report.Suites, *s)
		}

		log.Printf("DEBUG: REPORT %#v", report)
		return report, nil
	} else {
		log.Println("Unknown format.")
		return nil, errors.New("Unknown format.")
	}
}
