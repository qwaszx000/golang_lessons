package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func test_func1(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[%d] Start\n", id)

	sleep_timer := time.After(time.Millisecond * 15)
	value := ctx.Value(TEST_KEY)

	select {
	case <-sleep_timer:
	case <-ctx.Done():
		fmt.Printf("[%d] Context cancelled before timer!\n", id)
		return
	}

	fmt.Printf("[%d] Timer is out -- value=%d\n", id, value)
}

type ContextKey int

const TEST_KEY ContextKey = 0

func main() {
	wg := sync.WaitGroup{}
	ctx_root := context.Background()

	ctx1, cancel1 := context.WithCancel(ctx_root)
	defer cancel1()
	ctx1 = context.WithValue(ctx1, TEST_KEY, 5)

	//must timeout
	wg.Add(1)
	go test_func1(ctx1, 0, &wg)
	time.AfterFunc(time.Millisecond*5, func() {
		cancel1()
	})

	ctx2, cancel2 := context.WithTimeout(ctx_root, time.Millisecond*10)
	defer cancel2()
	ctx2 = context.WithValue(ctx2, TEST_KEY, 7)

	//also must timeout
	wg.Add(1)
	go test_func1(ctx2, 1, &wg)

	ctx3, cancel3 := context.WithCancel(ctx_root)
	defer cancel3()
	ctx3 = context.WithValue(ctx3, TEST_KEY, 8)

	//Must not timeout
	wg.Add(1)
	go test_func1(ctx3, 2, &wg)

	wg.Wait()
}
