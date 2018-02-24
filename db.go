package main

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
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
	Id    uint64 `gorm:"primary_key"`
	Name  string `gorm:index` // TAP, SubUnit, JUnit
	Tests []Test
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

func initDb(dbpath string) *gorm.DB {

	db, err := gorm.Open("sqlite3", dbpath)
	db.AutoMigrate(&Report{}, &Suite{}, &Test{})

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
