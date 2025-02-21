package main

import (
	"context"
	"testing"
	"testing/synctest" //requires env var: GOEXPERIMENT=synctest
)

func TestFuncToTest(t *testing.T) {
	synctest.Run(func() {
		ctx, cancel := context.WithCancel(context.Background())
		test_chan := make(chan int, 0)

		go FuncToTest(ctx, test_chan)

		synctest.Wait()

		select {
		case <-test_chan:
			cancel()
			t.Error("FuncToTest sent data to resp channel before it was expected\n")
			return
		default:
			break
		}

		cancel()

		synctest.Wait()

		select {
		case <-test_chan:
			break
		default:
			t.Error("FuncToTest haven't sent data to resp channel after cancel\n")
		}
	})
}

// Expected to fail
func TestFuncToTestBad(t *testing.T) {
	synctest.Run(func() {
		ctx, cancel := context.WithCancel(context.Background())
		test_chan := make(chan int, 0)

		go FuncToTestBad(ctx, test_chan)

		synctest.Wait()

		select {
		case <-test_chan:
			cancel() //make bad gorutine return to prevent deadlock
			t.Error("FuncToTestBad sent data to resp channel before it was expected\n")
		default:
			break
		}

		cancel()

		synctest.Wait()

		select {
		case <-test_chan:
			break
		default:
			t.Error("FuncToTestBad haven't sent data to resp channel after cancel\n")
		}
	})
}

//Run as `GOEXPERIMENT=synctest go test`
