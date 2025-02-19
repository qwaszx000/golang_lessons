package main

import (
	"fmt"
)

type AnyType interface{}
type Number interface {
	~int | float32 | float64
}

type EmptyStruct struct{}
type TreeNode[T AnyType] struct {
	left, right *TreeNode[T]
	value       T
}

func (node *TreeNode[T]) print_node() {
	fmt.Println(node.value)

	/*if _, ok := node.value.(Number); ok {
		fmt.Println("Is Num")
	}

	switch node.value.(type) {
	case int:
		fmt.Println("Is int")
	case string:
		fmt.Println("Is string")
	}*/ //doesn't works and errors because it's go's wart

	//fix with any()
	/*if _, ok := any(node.value).(Number); ok {
		fmt.Println("Is Num")
	}*/ //Still errors, because Number can only be used as type constraint
	//var aaa Number = 0 //can't do this because of the same reason

	switch any(node.value).(type) {
	case int:
		fmt.Println("Is int")
	case string:
		fmt.Println("Is string")
	}
}

func main() {

	//Type assertions
	var test1 AnyType = "1234"

	value, ok := test1.(string)
	fmt.Println(test1, value, ok)

	value2, ok := test1.(int)
	fmt.Println(test1, value2, ok)

	/*v := test1.(float32) //panics because we expect float32 without second argument to test if it's true
	fmt.Println(v)*/

	//type switch1
	switch test1.(type) {
	case int:
		fmt.Println("Is int")
	case AnyType:
		fmt.Println("Is AnyType")
	case string:
		fmt.Println("Is string")
	default:
		break
	}

	//type switch2
	switch test1.(type) {
	case int:
		fmt.Println("Is int")
	case string:
		fmt.Println("Is string")
	case AnyType:
		fmt.Println("Is AnyType")
	default:
		break
	}

	//Generics
	var node1 TreeNode[int]
	node1.value = 5
	node1.print_node()
}
