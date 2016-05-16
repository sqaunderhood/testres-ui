package main

import (
	"fmt"
	"strings"
)

type Token int

const (
	TOKEN_IN Token = iota
	TOKEN_FORMAT
	TOKEN_STATUS
	TOKEN_CREATED
)

var tokens = [...]string{
	"in",
	"format",
	"status",
	"created",
}

func (t Token) String() string { return tokens[t-1] }

func makequery() {

	const src = `cat format:junit status:pass created:2006-08-10`

	lexemes := strings.Fields(src)
	for _, w := range lexemes {
		value := strings.Split(w, ":")
		if len(value) == 1 {
			fmt.Println("This is value without keyword ", value[0])
			continue
		}

		value[0] = strings.ToLower(value[0])
		switch value[0] {
		case "in":
			fmt.Println("search in specific field:", value[1])
		case "format":
			fmt.Println("search by format:", value[1])
		case "status":
			fmt.Println("search by status:", value[1])
		case "created":
			fmt.Println("search by date:", value[1])
		}
	}
}

func search(query string) []Report {

	db := initDb(dbpath)
	defer db.Close()

	var reports []Report
	db.Where("body = ?", query).Find(&reports)

	return reports
}
