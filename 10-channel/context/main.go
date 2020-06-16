package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("number of goroutine : ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				fmt.Println("woking on n: ", n)
			}
		}
	}()

	time.Sleep(time.Millisecond * 300) // wait until 300 milis
	fmt.Println("number of goroutine : ", runtime.NumGoroutine())

	cancel()
	time.Sleep(time.Second)
	fmt.Println("number of goroutine : ", runtime.NumGoroutine())
}
