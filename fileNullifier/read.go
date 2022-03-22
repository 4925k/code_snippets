package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Setting logging

	log.Println("open file")
	f, err := os.OpenFile("nullify.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600) // This will be our logging file
	if err != nil {
		log.Println("[ERROR] opening file", err)
		return
	}
	log.SetFlags(log.Lshortfile)
	log.SetOutput(f)

	log.Println("testing")

	//get stats from the file
	stats, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	n := stats.Size()

	// var temp [int(n)]byte
	temp := make([]byte, n)
	err = os.WriteFile("test.txt", temp, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
