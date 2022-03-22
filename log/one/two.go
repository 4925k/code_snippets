package one

import "fmt"

func Two() error {
	return fmt.Errorf("two.go: %v", Ek())
}
