package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var a uint64 = 0

	for index := 0; index < 100; index++ {
		go func() {
			for {
				atomic.AddUint64(&a, 2)

				runtime.Gosched()
			}

		}()
	}

	time.Sleep(time.Second * 2)
	b := atomic.LoadUint64(&a)
	fmt.Print(b)
}
