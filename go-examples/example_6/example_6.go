package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var Cishu uint64 = 0
	var status = make(map[int]int, 0)
	var l = new(sync.Mutex)

	for a := 0; a < 100; a++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				l.Lock()
				total += status[key]
				l.Unlock()
				atomic.AddUint64(&Cishu, 1)
				runtime.Gosched()
			}
		}()
	}

	for b := 0; b < 10; b++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				l.Lock()
				status[key] = val
				l.Unlock()
				atomic.AddUint64(&Cishu, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	rcishu := atomic.LoadUint64(&Cishu)
	fmt.Println(rcishu)
	l.Lock()
	fmt.Println(status)
	l.Unlock()
}
