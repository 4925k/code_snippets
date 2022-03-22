// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

var value interface{}

func main() {

	value = []string{"1.1.1.1", "2.2.2.2"}
	temp := reflect.ValueOf(value)
	tempV := []string{}

	fmt.Printf("Type: %T,\nValue:%v\n", temp, temp)
	fmt.Printf("Type: %T,\nValue:%v\n", tempV, tempV)

	for i := 0; i < temp.Len(); i++ {
		tempV = append(tempV, temp.Index(i).Interface().(string)) // currently only slice of string
	}

	fmt.Printf("Type: %T,\nValue:%v\n", temp, temp)
	fmt.Printf("Type: %T,\nValue:%v\n", tempV, tempV)
}
