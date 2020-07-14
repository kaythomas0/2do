package main

import (
	"log"

	"github.com/namsral/flag"
	"github.com/prologic/bitcask"
)

var (
	db *bitcask.Bitcask
)

func main() {
	var (
		dbpath string
		bind   string
	)

	flag.StringVar(&dbpath, "dbpath", "todo.db", "Database path")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.Parse()

	var err error
	db, err = bitcask.Open("/2do/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	newServer(bind).listenAndServe()
}
