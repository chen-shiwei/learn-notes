package main

import (
	"fmt"
	"time"
)

func main() {

	// var requests = make(chan int, 5)
	// for i := 0; i < 5; i++ {
	// 	requests <- i
	// }
	// var ticker = time.NewTicker(time.Millisecond * 200)

	// for t := range ticker.C {
	// 	fmt.Println(t, <-requests)
	// }

	var burstyLimiter = make(chan time.Time, 5)

	for a := 0; a < 4; a++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	var requests = make(chan int, 5)
	for b := 0; b < 5; b++ {
		requests <- b
	}
	close(requests)
	for req := range requests {
		<-burstyLimiter
		fmt.Println(req, time.Now())
	}
}
