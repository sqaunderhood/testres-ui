package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Report struct {
	gorm.Model
	Format   Format `db:"format"`
	ReportId string `db:"reportid"`
	Filename string `db:"filename"`
	Body     string `db:"body"`
	Hits     int    `db:"hits"`
	Suites   []*Suite
}

type Suite struct {
	gorm.Model
	Name  string  `db:"size:255;index:name_idx"` // TAP, SubUnit, JUnit
	Tests []*Test // TAP, SubUnit, JUnit
}

type Test struct {
	gorm.Model
	Name        string `db:"name"`        // SubUnit, JUnit
	Status      Status `db:"status"`      // TAP, SubUnit, JUnit
	Ok          bool   `db:"ok"`          // TAP, SubUnit
	Description string `db:"description"` // TAP
	Explanation string `db:"explanation"` // TAP
	StartTime   string `db:"starttime"`   // SubUnit, JUnit
	EndTime     string `db:"endtime"`     // SubUnit, JUnit
	Tags        string `db:"tags"`        // SubUnit
	Details     []byte `db:"details"`     // TAP, SubUnit, JUnit
}

func initDb(dbpath string) *gorm.DB {

	db, err := gorm.Open("sqlite3", dbpath)

	if err != nil {
		log.Println("gorm.Open failed")
		return nil
	}

	if os.Getenv("DEBUG") == "true" {
		log.Println("Debug mode enabled")
		return db.Debug().LogMode(true)
	}

	if !db.HasTable(Report{}) {
		db.CreateTable(&Report{})
	}
	if !db.HasTable(Suite{}) {
		db.CreateTable(&Suite{})
	}
	if !db.HasTable(Test{}) {
		db.CreateTable(&Test{})
	}

	return db
}
