package main

import (
	"database/sql"
	"fmt"
)

type suck struct {
	item  string
	count int
	db    *sql.DB
}

func main() {
	var s suck
	x := suck{}
	y := new(suck)

	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", x)
	fmt.Printf("%+v\n", y)

}
