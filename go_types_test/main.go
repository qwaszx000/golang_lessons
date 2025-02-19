package main

import (
	"fmt"
	"package_test/types"
)

func Test() {
	fmt.Println("Test")
}

func main() {
	struct1 := new(types.MyExampleStruct)
	struct2 := types.MyExampleStruct{}

	var interface_obj types.DoSmthInterface

	interface_obj = struct1 //copies pointer, changes are seen
	interface_obj.DoSmth()

	//struct1.func_ptr = &Test //gives error
	var some_int int = 5

	struct1.Func_simple = Test
	struct1.Num = 5
	struct1.Str = "Test\tTest"
	struct1.Int_ptr = &some_int
	struct1.My_slice = make([]byte, 2, 4)
	struct1.Is_true = true
	struct1.My_array = [5]int{0, 1, 2}
	struct1.My_map = make(map[string]int, 2)
	struct1.My_map["a"] = 1
	struct1.My_map["b"] = 2

	interface_obj.DoSmth() //same as struct1
	struct1.DoSmth()

	//-------
	interface_obj = struct2 //copies by value, so next changes are not accesible from interface_obj
	interface_obj.DoSmth()

	//struct1.func_ptr = &Test //gives error
	struct2.Func_simple = Test
	struct2.Num = 5
	struct2.Str = "Test\tTest"
	struct2.Int_ptr = &some_int
	struct2.My_slice = make([]byte, 2, 4)
	struct2.Is_true = true
	struct2.My_array = [5]int{0, 1, 2}
	struct2.My_map = make(map[string]int, 2)
	struct2.My_map["a"] = 1
	struct2.My_map["b"] = 2
	//struct2.hidden_internal_num = 5 //error because it is not exported from package types

	interface_obj.DoSmth() //still empty
	struct2.DoSmth()       //changes are seen

	//--------
	var my_test_struct5 types.Struct2 = *new(types.Struct2)
	my_test_struct5.Aaa = 5
	my_test_struct5.Bbb = 6
	fmt.Println(my_test_struct5)
}
