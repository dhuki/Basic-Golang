package main

func main() {
	// go specialize built for concurrency to getting advantage for multiprocessor (dual core, core i3, etc) nowayday
	// and makes a huge significant different when do concurrency with other language
	// concurency in other language seem difficult and strugle to implement
	// otherwise in go it so easy to do conccurency

	// concurrency is not parallelism
	// concurrency is the way that you make something or write a code (in this case)
	// to be able possibility run in parallel / in same time (we can just called it DESIGN PATTERN)
	// but concurrency not guaranteed your code will run parallelism bcs the primary reason
	// is you have to run in multiple CPU cores (brain of your computer, core i5 usually has 4 CPU's cores)
	// ROB PIKE CONCURRENCY IS NOT PARALLELISM : ON YOUTUBE

	// wachout about race condition
	// race condition in conccurent programming is when you have shared variable between two goroutine
	// and screw up by read or write and making variable not consistent
	// CHECK ARDANLAB DATA RACE
}
