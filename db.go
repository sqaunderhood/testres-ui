package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ligurio/recidive/formats"
	_ "github.com/mattn/go-sqlite3"
)

func initDb(dbpath string) *gorm.DB {

	db, err := gorm.Open("sqlite3", dbpath)
	db.AutoMigrate(&formats.Report{}, &formats.Suite{}, &formats.Test{})

	if err != nil {
		log.Println("gorm.Open failed")
		return nil
	}

	if os.Getenv("DEBUG") == "true" {
		log.Println("Debug mode enabled")
		return db.Debug().LogMode(true)
	}

	if !db.HasTable(formats.Report{}) {
		db.CreateTable(&formats.Report{})
	}
	if !db.HasTable(formats.Suite{}) {
		db.CreateTable(&formats.Suite{})
	}
	db.Model(&formats.Report{}).Related(&formats.Suite{}, "ReportId")
	if !db.HasTable(formats.Test{}) {
		db.CreateTable(&formats.Test{})
	}
	db.Model(&formats.Suite{}).Related(&formats.Test{}, "SuiteId")

	return db
}
