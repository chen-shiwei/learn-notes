package main

//extern void SayHello(_GoString_ s);
import "C"
import "fmt"

//export SayHello
func SayHello(s string) {
	fmt.Println(s)
}

func main() {
	C.SayHello("c go!")
}
