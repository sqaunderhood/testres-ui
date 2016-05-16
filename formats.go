package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/ligurio/go-junit/parser"
	"github.com/ligurio/go-subunit/parser"
	"github.com/ligurio/go-tap/parser"
)

// Supported formats
var Formats = [...]string{"SubUnit", "JUnit", "TAP"}

type Report struct {
	Id       int64  `db:"id"`
	Format   string `db:"format"`    // name of report format
	ReportId string `db:"reportid"`  // report id is uniq identificator
	Filename string `db:"filename"`  // file name of report
	Body     []byte `db:"body"`      // raw report
	Created  int64  `db:"timestamp"` // timestamp when report was uploaded
	Hits     int    `db:"hits"`
	Suites   []*Suite
}

type Suite struct {
	Id    int64   `db:"id"`
	Name  string  `db:"name"` // TAP, SubUnit, JUnit
	Tests []*Test // TAP, SubUnit, JUnit
}

type Test struct {
	Id          int64         `db:"id"`
	Name        string        `db:"name"`        // SubUnit, JUnit
	Status      tap.Directive `db:"status"`      // TAP, SubUnit, JUnit
	Ok          bool          `db:"ok"`          // TAP, SubUnit
	Description string        `db:"description"` // TAP
	Explanation string        `db:"explanation"` // TAP
	StartTime   string        `db:"starttime"`   // SubUnit, JUnit
	EndTime     string        `db:"endtime"`     // SubUnit, JUnit
	Tags        []string      `db:"tags"`        // SubUnit
	Details     []byte        `db:"details"`     // TAP, SubUnit, JUnit
}

type Status int

const (
	None Status = iota
	Fail
	Pass
	Skip
	Error
	Todo
	XFail
	UxSuccess
)

func (s Status) String() string {
	switch s {
	case None:
		return "None"
	case Fail:
		return "Fail"
	case Pass:
		return "Pass"
	case Skip:
		return "Skip"
	case Error:
		return "Error"
	case Todo:
		return "Todo"
	case XFail:
		return "XFail"
	case UxSuccess:
		return "UxSuccess"
	}
	return ""
}

func ReadReport(r io.Reader, name string) (error, *Report) {

	report := new(Report)

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err, nil
	}

	report.Body = buf

	report.ReportId = makeid()
	log.Println("Report ID is", report.ReportId)

	report.Filename = name
	report.Created = time.Now().UnixNano()
	_, err = tap.NewParser(r)
	log.Println("error:", err)

	if p, err := tap.NewParser(r); err == nil {
		report.Format = Formats[2]
		log.Println(report.Format)
		ts, _ := p.Suite()
		if err != nil {
			log.Println("error reading suites", err)
		}

		s := new(Suite)
		s.Name = ""

		for _, tl := range ts.Tests {
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
		return nil, report

	} else if err, _ := junit.Parser(r); err == nil {
		log.Println("JUnit format detected.")
		report.Format = Formats[1]

		err, p := junit.Parser(r)
		if err != nil {
			log.Println("error", err)
			return err, nil
		}

		for _, ts := range p.Testsuites {
			s := new(Suite)
			for _, test := range ts.Testcases {
				t := new(Test)
				t.Name = test.Name
				// FIXME t.Status = test.Status
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

		log.Printf("%#v", report)
		return nil, report

	} else if _, err := subunit.Parser(r); err == nil {
		log.Println("SubUnit format detected.")
		report.Format = Formats[0]

		p, err := subunit.Parser(r)
		if err != nil {
			log.Println("error", err)
			return err, nil
		}
		s := new(Suite)
		s.Name = p.Test

		for _, test := range p.Tests {
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

		log.Printf("%#v", report)
		report.Suites = append(report.Suites, s)
		return nil, report

	} else {
		log.Println("Unknown format.")
		return errors.New("Unknown format."), nil

	}
}
