package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// use atomic, atomic is antoher way to prevent data race

	var wg sync.WaitGroup

	var counter int64 // when use int64 think about using atomic
	length := 100
	wg.Add(length)

	for i := 0; i < length; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // allowing go routine to run
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.NumGoroutine()
			wg.Done()
		}()
	}
	wg.Wait()
	// counter not consistent bcs race condition happen among goroutine
	// try to run (go run -race main.go) to check race condition
	fmt.Println(counter)

}
