package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	length := 100

	wg.Add(length)

	for i := 0; i < length; i++ {
		go func() {
			v := counter
			runtime.Gosched() // allowing go routine to run
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	// counter not consistent bcs race condition happen among goroutine
	// try to run (go run -race main.go) to check race condition
	fmt.Println(counter)
}
