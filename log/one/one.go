package one

import (
	"fmt"
	"os"
)

func Ek() error {
	_, err := os.OpenFile("asd.txt", os.O_APPEND, 0666)
	ceLogger.log("asd")
	return fmt.Errorf("one.go: %v", err)
}
