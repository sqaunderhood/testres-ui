package main

import (
	"flag"
	"log"
)

const dbpath = "db.sqlite"

var (
	httpAddr  = flag.String("http", ":8080", "HTTP service address and port")
	staticDir = flag.String("static", "./static/", "static and templates files directory")
)

func main() {
	flag.Parse()
	log.Fatal(StartServer(*httpAddr, staticDir))
}
