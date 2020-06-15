package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// print system computer
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU cores\t", runtime.NumCPU()) // cpu cores is like brain on your computer
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	// it will exited when main goroutine is done without wait another go routine done
	// go foo()
	// bar()

	// we can sync gorountine wait until all done and exited
	// by using mutex or wait group
	// this use group
	wg.Add(1) // add one other go routine
	go foo()
	bar()
	wg.Wait() // wait until all done
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
