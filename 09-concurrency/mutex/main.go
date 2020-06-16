package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// use mutex to prevent data race
	// mutex basically is locking variable or chuck of code that use in goroutine so another goroutine cannot access that variable
	// until mutex unlock the process

	var wg sync.WaitGroup

	counter := 0
	length := 100

	wg.Add(length)

	var mutext sync.Mutex

	for i := 0; i < length; i++ {
		go func() {
			mutext.Lock() // will lock variable below from another goroutine that want use
			v := counter
			runtime.Gosched() // allowing go routine to run
			v++
			counter = v
			mutext.Unlock() // release lock so other goroutine can use
			wg.Done()
		}()
	}
	wg.Wait()
	// counter not consistent bcs race condition happen among goroutine
	// try to run (go run -race main.go) to check race condition
	fmt.Println(counter)

	// there is another use of mutex that is RWMutex, it's more flexibel than mutex
}
