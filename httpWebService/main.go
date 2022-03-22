package main

import (
	"database/sql"
	"log"
)

type server struct {
	db     *sql.DB
	router *server
	log    *log.Logger
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("hello world")
}
