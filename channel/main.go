package main

import "fmt"

func main() {
	EndSignal := make(chan interface{}, 1)
	select {
	case <-EndSignal: // If the end goroutine signal was sent
		fmt.Println("here")
	default:
		fmt.Println("here")
	}

	fmt.Println("end")
}
