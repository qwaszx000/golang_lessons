package main

import (
	"fmt"
	"testing/synctest" //requires env var: GOEXPERIMENT=synctest
)

// TODO
func main() {
	fmt.Println("Start")

	_ = synctest.Wait
}
