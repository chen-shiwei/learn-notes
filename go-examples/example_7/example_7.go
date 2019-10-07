package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type reader struct {
	key  int
	resp chan int
}

type writer struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readers = make(chan *reader)
	var writers = make(chan *writer)

	var count uint64 = 0

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case readAction := <-readers:
				readAction.resp <- state[readAction.key]
			case writeAction := <-writers:
				state[writeAction.key] = writeAction.val
				writeAction.resp <- true
			}
		}
	}()

	for a := 0; a < 100; a++ {
		go func() {
			for {
				r := &reader{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				readers <- r
				<-r.resp
				atomic.AddUint64(&count, 1)
			}
		}()
	}

	for b := 0; b < 10; b++ {
		go func() {
			for {
				w := &writer{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writers <- w
				<-w.resp
				atomic.AddUint64(&count, 1)
			}
		}()
	}
	time.Sleep(time.Second)

	B := atomic.LoadUint64(&count)

	fmt.Print(B)

}
