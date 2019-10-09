package main

//#include <stdio.h>
// extern void SayHello(char* s);
import "C"
import "fmt"

//export SayHello
func SayHello(s *C.char) {
	fmt.Println(C.GoString(s))
}

func main() {
	C.SayHello(C.CString("c go!"))
}
