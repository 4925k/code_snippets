package main

import (
	"fmt"
	"time"
)

var POSTGRES_TIMESTAMP_LAYOUT = "2006-01-02 15:04:05.999999999Z07:00"

func main() {
	timestamp := time.Now().Format(POSTGRES_TIMESTAMP_LAYOUT)
	fmt.Println(timestamp)

	local := time.Now().Local().Format(POSTGRES_TIMESTAMP_LAYOUT)
	fmt.Println(local)
}
