package main

import (
	"context"
	"fmt"
)

func FuncToTest(ctx context.Context, resp chan<- int) {
	//Wait for Done to close
	<-ctx.Done()

	//Send responce
	resp <- 5
}

func FuncToTestBad(ctx context.Context, resp chan<- int) {
	//Send responce
	resp <- 5

	//Wait for Done to close
	<-ctx.Done()
}

func main() {
	fmt.Println("Start")
}
