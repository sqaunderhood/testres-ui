package formats

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/ligurio/recidive/formats/junit"
)

type Report struct {
	Id        uint64 `gorm:"primary_key"`
	UID       string
	Format    string
	Filename  string
	Body      string
	Hits      int
	Suites    []Suite
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Suite struct {
	Id       uint64 `gorm:"primary_key"`
	ReportId uint64
	Name     string `gorm:index` // TAP, SubUnit, JUnit
	Tests    []Test
}

type Test struct {
	Id          uint64 `gorm:"primary_key"`
	SuiteId     uint64
	Name        string `gorm:index` // SubUnit, JUnit
	Ok          bool   // TAP, SubUnit
	Description string // TAP
	Explanation string // TAP
	StartTime   string // SubUnit, JUnit
	EndTime     string // SubUnit, JUnit
	Tags        string // SubUnit
	Details     []byte // TAP, SubUnit, JUnit
	Status      Status // TAP, SubUnit, JUnit
}

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
