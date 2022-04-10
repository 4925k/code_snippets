package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("insufficient arguments")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open file: %v", err)
	}

	scanner := bufio.NewScanner(f)

	count := 0

	for scanner.Scan() {
		count++
	}

	fmt.Println(count)
}
