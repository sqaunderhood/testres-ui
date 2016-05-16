package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func initDb(dbpath string) *gorm.DB {

	db, err := gorm.Open("sqlite3", dbpath)

	if err != nil {
		log.Println("gorm.Open failed")
		return nil
	}

	if !db.HasTable(Report{}) {
		db.CreateTable(Report{})
	}
	if !db.HasTable(Suite{}) {
		db.CreateTable(Suite{})
	}
	if !db.HasTable(Test{}) {
		db.CreateTable(Test{})
	}

	if os.Getenv("DEBUG") == "true" {
		log.Println("Debug mode enabled")
		return db.Debug().LogMode(true)
	}

	return db
}
