package main

import (
	"fmt"
)

// Team lead at one interview told me defer will print 5 in this case
// And it's true
func defer_test1() {
	var x int = 5
	defer fmt.Println(x) //prints 5, wtf???
	x = 6
	fmt.Println(x)
}

// What if we do it with pointer?
// Still true
func defer_test2() {
	var x int = 5
	var x_ptr *int = &x
	defer fmt.Println(*x_ptr) //prints 5, wtf??? x2
	*x_ptr = 6
	fmt.Println(*x_ptr)
}

// Also he told me defers are executed in reverse order
// Yep, also true
func defer_order_test() {
	defer fmt.Println("")
	for x := 0; x < 5; x++ {
		defer fmt.Printf("%d", x)
	}
}

// Another question was about changing slice elements by append and by index access
// Let's test it too
func _slice_test1(slice_inp []int) {
	slice_inp = append(slice_inp, 5) // So append is in functional style; ok, didn't know that
	slice_inp[0] = 5                 // It changed real slice, as lead said; i though you need to pass pointer
	// But it really looks like slice is just array with metadata
	// Something like this:
	/*type Slice struct {
		array_ptr *[n]T
		start_index int
		len uint
	}*/
	// It means when we pass slice by value, pointer still points to the same array

	fmt.Println("Internal slice state now: ", slice_inp)
}
func slice_test1() {
	var test_slice []int = make([]int, 2, 4)
	test_slice[0], test_slice[1] = 1, 2
	fmt.Println("Initial slice state: ", test_slice)

	_slice_test1(test_slice)

	fmt.Println("Slice state now: ", test_slice)
}

// Second level of this question was about relocating of slice when we exceeed it's capacity
// I expect it to change internal array addr because of doing malloc
// So in result we should have same slice as before calling _slice_test2
func _slice_test2(slice_inp []int) {

	slice_inp = append(slice_inp, 5, 6, 7) // Append 3 elements to exceed capacity
	/*for i := len(slice_inp); i <= cap(slice_inp); i++ {
		slice_inp = append(slice_inp, i)
	}*/ //It was stupid, i admit; Result - infinite loop

	//Also he told me it will increase cap x2 when len and cap is small
	fmt.Println("New slice data: ", slice_inp, len(slice_inp), cap(slice_inp))
	//Yep, it does: new cap == 8

	slice_inp[0] = 5

	fmt.Println("Internal slice state now: ", slice_inp)
}
func slice_test2() {
	var test_slice []int = make([]int, 2, 4)
	test_slice[0], test_slice[1] = 1, 2
	fmt.Println("Initial slice state: ", test_slice)

	_slice_test2(test_slice)

	fmt.Println("Slice state now: ", test_slice)
	//I was right, it's still the same slice
}

func main() {

	/*var float_test float32 = 1.9
	var float_test2 float64 = 1.9

	fmt.Printf("%d %d %d %d\n", float_test, float_test2, int(float_test), int(float_test2))*/ //results in "%!d(float32=1.9) %!d(float64=1.9) 1 1"

	defer_test1()
	defer_test2()
	defer_order_test()
	slice_test1()
	slice_test2()
}
