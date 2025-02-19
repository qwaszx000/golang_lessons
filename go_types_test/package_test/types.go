package types

import "fmt"

type DoSmthInterface interface {
	DoSmth()
	DoSmth2(bool) int
}

type MyExampleStruct struct {
	Num                 int
	Str                 string
	Int_ptr             *int
	My_slice            []byte
	Is_true             bool
	My_array            [5]int
	My_map              map[string]int
	Func_ptr            *func()
	Func_simple         func()
	hidden_internal_num int
}

func (this MyExampleStruct) DoSmth() {
	fmt.Println(this)
}

func (this MyExampleStruct) DoSmth2(inp bool) int {
	if inp {
		return 5
	} else {
		return 6
	}
}

type Struct1 struct {
	Aaa int
}

type Struct2 struct {
	Struct1
	Bbb int
}
