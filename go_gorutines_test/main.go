package main

import (
	"fmt"
	"sync"
)

func gorutine_test(list []int, c_allowed <-chan int, c_next chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()
	defer close(c_next)

	for _, val := range list {

		//wait for our turn
		<-c_allowed

		//do printing
		fmt.Println(val)

		//tell about next turn
		//causes error in the last step because we are sending signal and no one is listening if we make channels unbuffered
		//so we make them buffered by make(chan int, 1) instead of make(chan int, 0)
		c_next <- 1
	}
}

type MyMutexSwitch struct {
	lock_mutex          sync.Mutex
	current_gorutine_id int
}

const _GORUTINES_COUNT int = 2

func gorutine_test_mutex(list []int, go_id int, wg *sync.WaitGroup, mutex_switch *MyMutexSwitch) {

	defer wg.Done()

	for _, val := range list {
		mutex_switch.lock_mutex.Lock()
		for mutex_switch.current_gorutine_id != go_id {
			mutex_switch.lock_mutex.Unlock()
			mutex_switch.lock_mutex.Lock()
		}

		//do printing
		fmt.Println(val)

		if go_id >= _GORUTINES_COUNT-1 {
			mutex_switch.current_gorutine_id = 0
		} else {
			mutex_switch.current_gorutine_id = go_id + 1
		}

		mutex_switch.lock_mutex.Unlock()
	}
}

func gorutine_panic_recovery_test(y int, wg *sync.WaitGroup) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic-intended gorutine: ", err)
		}
	}()

	wg.Done()
	var x int = 5 / y //Intended -> panic: runtime error: integer divide by zero
	fmt.Println(x)
}

func main() {
	var wg sync.WaitGroup
	var lists_array [_GORUTINES_COUNT][]int = [_GORUTINES_COUNT][]int{
		[]int{1, 3, 5, 7, 9},
		[]int{2, 4, 6, 8, 10},
	}

	var mutex_switch MyMutexSwitch = MyMutexSwitch{}
	mutex_switch.current_gorutine_id = 0

	//mutex way
	for i := 0; i < _GORUTINES_COUNT; i++ {
		wg.Add(1)
		go gorutine_test_mutex(lists_array[i], i, &wg, &mutex_switch)
	}
	wg.Wait()

	//channel way
	var communication_channels [_GORUTINES_COUNT]chan int
	for i, _ := range communication_channels {
		communication_channels[i] = make(chan int, 1) //buffered to prevent deadlock on last channel send operation
	}

	for i := 0; i < _GORUTINES_COUNT; i++ {
		wg.Add(1)

		//send next to current i
		//receive allowed from i-1 or last element

		//first gorutine receives allowed from last gorutine
		if i == 0 {
			go gorutine_test(lists_array[i], communication_channels[_GORUTINES_COUNT-1], communication_channels[0], &wg)
		} else {
			go gorutine_test(lists_array[i], communication_channels[i-1], communication_channels[i], &wg)
		}
	}
	communication_channels[_GORUTINES_COUNT-1] <- 1

	wg.Wait()

	//panic - recovery
	wg.Add(1)
	go gorutine_panic_recovery_test(0, &wg)
	wg.Wait()

	fmt.Println("main is still alive")
}
