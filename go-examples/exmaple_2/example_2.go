package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C
		fmt.Print("时间到")
	}()

	time.Sleep(time.Second * 3)
}
