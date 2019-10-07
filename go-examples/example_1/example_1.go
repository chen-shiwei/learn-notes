package main

import (
	"fmt"
	"time"
)

func main() {
	var a = make(chan string, 0)
	go func() {
		time.Sleep(time.Second * 3)
		a <- "result"
	}()

	select {
	case result := <-a:
		fmt.Printf("receive :%s", result)
	case <-time.After(time.Second * 4):
		fmt.Print("已超时")

	}
}
