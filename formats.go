package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	_ "time"

	"github.com/ligurio/go-junit/parser"
	_ "github.com/ligurio/go-subunit/parser"
	_ "github.com/ligurio/go-tap/parser"
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
	report.Body = buf

	r = bytes.NewReader(buf)

	report.ReportId = makeid()
	log.Println("Report ID is", report.ReportId)

	report.Filename = name
	//report.Created = time.Now().UnixNano()

	if jreport, err := junit.NewParser(r); err == nil {
		log.Println("DEBUG: JUnit format detected.")
		//report.Format = FmtJUnit

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
				s.Tests = append(s.Tests, t)
			}
			report.Suites = append(report.Suites, s)
		}

		log.Printf("DEBUG: REPORT %#v", report)
		return report, nil

		/*
			} else if treport, err := tap.NewParser(r); err != nil {
				report.Format = FmtTAP
				log.Println("DEBUG: TAP format detected.")

				ts, err := treport.Suite()
				if err != nil {
					log.Println("error reading TAP suites", err)
				}

				s := new(Suite)
				s.Name = ""

				for _, tl := range treport.Tests {
					t := new(Test)
					//t.Name = ""
					t.Status = tl.Directive
					t.Ok = tl.Ok
					t.Description = tl.Description
					t.Explanation = tl.Explanation
					//t.StartTime = ""
					//t.EndTime = ""
					//t.Tags = ""
					t.Details = tl.Yaml
					s.Tests = append(s.Tests, t)
				}

				report.Suites = append(report.Suites, s)

				log.Printf("%#v", report)
				return report, nil

			} else if sreport, err := subunit.NewParser(r); err == nil {
				log.Println("SubUnit format detected.")
				report.Format = FmtSubUnit

				s := new(Suite)
				s.Name = sreport.Test

				for _, test := range sreport.Tests {
					t := new(Test)
					//t.Name = ""
					//FIXME t.Status = test.State
					//t.Ok = ""
					t.Description = test.Label
					//t.Explanation = ""
					//t.StartTime = ""
					//t.EndTime = ""
					t.Tags = test.Tags
					//t.Details = ""
					s.Tests = append(s.Tests, t)
				}

				log.Printf("REPORT %#v", report)
				report.Suites = append(report.Suites, s)
				return report, nil
		*/
	} else {
		log.Println("Unknown format.")
		return nil, errors.New("Unknown format.")
	}
}
