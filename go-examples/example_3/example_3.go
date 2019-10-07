package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)

	go func() {
		for t := range ticker.C {
			fmt.Print(t)
		}
	}()

	time.Sleep(time.Second * 3)
}
