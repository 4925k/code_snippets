package main

import (
	"log"
	"logTesting/one"
	"os"
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ceLogger := log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func main() {

	ceLogger.Println("error from main")
	ceLogger.Print(one.Ek())
	ceLogger.Print(one.Two())
}
