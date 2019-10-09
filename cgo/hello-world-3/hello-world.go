package main

import "C"

//#include "./hello.c"
//#include <stdio.h>
import "C"

func main() {
	C.SayHello(C.CString("你好，cgo"))
}
