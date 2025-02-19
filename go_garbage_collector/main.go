package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// Also we talked about ability to organize struct fields in order to optimize it's size
// So this struct is not optimized
// Total size == 40
type Block struct {
	a int64  //8 bytes
	b uint64 //8 bytes
	c int8   //1 byte + 7 bytes padding
	d int64  //8 bytes
	e int8   //1 byte + 7 bytes padding
}

// And this struct is optimized
// Total size == 32
type OptimizedBlock struct {
	a int64  //8 bytes
	b uint64 //8 bytes
	d int64  //8 bytes
	e int8   //1 byte
	c int8   //1 byte + 6 bytes
}

func main() {

	fmt.Printf("Go version: %s\n", runtime.Version()) //go1.24.0

	b1 := Block{}
	b2 := OptimizedBlock{}

	fmt.Printf("Unoptimized block size: %d\nOptimized: %d\n", unsafe.Sizeof(b1), unsafe.Sizeof(b2))

	//Next we'll test GC
	runtime.GC()

	//Well, i can call it
	//But i don't see any way to set GOGC through code; i expect it to be env var for `go build`
}
