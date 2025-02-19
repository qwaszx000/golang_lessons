package main

import (
	"fmt"
	_ "strings" //import for side effects
)

func main() {
	save_to_cache("test1", []byte("test_val"))

	val := get_from_cache("test1")
	fmt.Println(string(val))

	save_to_db("test23", "test2_val3")
	val2 := get_from_db("test23")

	fmt.Println(val2)
	//fmt.Println("start")
	var _ = fmt.Println // to prevent unused import removing
	start_http_server()
}
