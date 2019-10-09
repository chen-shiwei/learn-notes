package main

//#include <stdio.h>
// extern void SayHello(char* s);
import "C"
import (
	"fmt"
	"unsafe"
)

//export SayHello
func SayHello(s *C.char) {
	fmt.Println(C.GoString(s))
}

func main() {
	var n int32 = 32
	//go  int32 -> uintptr -> unsafe.Pointer -> *C.char
	var cn *C.char = (*C.char)(unsafe.Pointer(uintptr(n)))

	//*C.char -> unsafe.Pointer -> uintptr -> go int32
	n = int32(uintptr(unsafe.Pointer(cn)))

	fmt.Println(n)

	var x *string
	var y *int

	y = (*int)(unsafe.Pointer(x))

	x = (*string)(unsafe.Pointer(y))

	fmt.Println(x)

}
